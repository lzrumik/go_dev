package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
	"sync"
	"time"
)

func getDate(datetime string) string {
	//datetime := "2015-01-01 00:00:00.0"  //待转化为时间戳的字符串

	return datetime[:10]

	timeLayout := "2006-01-02 15:04:05"  //转化所需模板
	loc, _ := time.LoadLocation("Local") //获取时区
	tmp, _ := time.ParseInLocation(timeLayout, datetime, loc)
	timestamp := tmp.Unix() //转化为时间戳 类型是int64

	data := time.Unix(timestamp, 0).Format("2006-01-02")
	//fmt.Println(time.Unix(timestamp, 0).Format("2006-01-02"))
	return data
}

var wg sync.WaitGroup
var ch = make(chan string, 10000)

func DealFile(filePath string, loadPath string,typeMap map[int]int){
	//打开文件
	file, _ := os.Open(filePath)
	defer file.Close()
	os.Truncate(loadPath, 0)

	//创建文件的缓冲读取器
	reader := bufio.NewReaderSize(file,4096*1024)
	line := 0

	for i := 0; i < 10; i++ {
		wg.Add(1)
		/*协程任务：从管道中拉取数据并写入到文件中*/
		go func(indx int) {
			//loadFilePathNew := loadPath + "-" + strconv.Itoa(indx) + ".txt"
			loadFilePathNew := loadPath
			loadFileHandle, err := os.OpenFile(loadFilePathNew, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}
			defer loadFileHandle.Close()

			//writer := bufio.NewWriter(loadFileHandle)
			writer := bufio.NewWriter(loadFileHandle)

			//totalLines := 0

			for lineStr := range ch  {
				writer.WriteString(lineStr)
			}
			writer.Flush()

			wg.Done()
		}(i)
	}

	//buffer := bytes.Buffer{}
	var buffer strings.Builder
	for {
		line++

		//逐行读取
		lineBytes, _, err := reader.ReadLine()

		//读到文件末尾时退出循环
		if err == io.EOF {
			break
		}

		//获取当前行的字符串

		//获取当前行的字符串
		//lineStr := string(lineBytes)

		repayInfo := new(RepayInfo)
		err2 := repayInfo.UnmarshalJSON(lineBytes)
		if err2 != nil {
			fmt.Println(err2)
			return
		}

		uid := repayInfo.UID
		repay_date := ""
		invest_amount := 0.00
		//fund_type := int8(repayInfo.CaCd)
		loan_id := repayInfo.LoanID

		if repayInfo.InvestRepayStatus == "REPAYED" {
			repay_date = getDate(repayInfo.InvestRepayTm)
			invest_amount = repayInfo.InvestRepayAmount
		}
		if repayInfo.InvestRepayStatus == "UNDUE" || repayInfo.InvestRepayStatus == "OVERDUE" {
			repay_date = getDate(repayInfo.InvestRepayDueDate)
			invest_amount = repayInfo.PlanRepayAmount
		}

		fund_type ,ok := typeMap[repayInfo.CaCd]
		if ok == false {
			fund_type = 0
		}

		buffer.WriteString(strconv.Itoa(uid))
		buffer.WriteString("\t")
		buffer.WriteString(repay_date)
		buffer.WriteString("\t")
		buffer.WriteString(strconv.FormatFloat(invest_amount, 'f', 2, 64))
		buffer.WriteString("\t")
		buffer.WriteString(strconv.Itoa(fund_type))
		buffer.WriteString("\t")
		buffer.WriteString(loan_id)
		buffer.WriteString("\n")

		if line%1 == 0 {
			ch <- buffer.String()
			buffer.Reset()
			//buffer = bytes.Buffer{}
		}
	}
	if buffer.String() != "" {
		ch <- buffer.String()
	}
	close(ch)

	wg.Wait()
	//wg.Done()
}

func main() {
	//创建输出文件
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile:", err)
	}

	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start cpu profile", err)
	}

	defer pprof.StopCPUProfile()

	//wg.Add(1)
	filePath := "/export/data/rsync/fund/app_crm_user_fix_repay_attribute_det.txt"
	loadPath := "/export/data/rsync/fund/app_crm_user_fix_repay_attribute_det.txt.load"

	//filePath = `D:\workspace_go\src\go_classify\repay\a.txt`
	//loadPath = `D:\workspace_go\src\go_classify\repay\a.load`

	typeMap := map[int]int{1:1,2:2}
	DealFile(filePath, loadPath,typeMap)

	f1, err := os.Create("mem.prof")
	if err != nil {
		log.Fatal("could not create mem profile:", err)
	}

	if err := pprof.WriteHeapProfile(f1); err != nil {
		log.Fatal("could not start mem profile", err)
	}
	f1.Close()

	f2, err := os.Create("goroutine.prof")
	if err != nil {
		log.Fatal("could not create goroutine profile:", err)
	}

	if gProf := pprof.Lookup("goroutine"); gProf == nil {
		log.Fatal("could not start goroutine profile", err)
	} else {
		gProf.WriteTo(f2, 0)
	}
	f2.Close()
	//wg.Add(1)
	//filePath = "/export/data/rsync/fund/app_crm_user_p2p_repay_attribute_det.txt"
	//loadPath = "/export/data/rsync/fund/app_crm_user_fix_repay_attribute_det.txt.load2"
	//go DealFile(filePath,loadPath)
	//
	//
	//wg.Add(1)
	//filePath = "/export/data/rsync/fund/app_crm_fengchu_repay_det.txt"
	//loadPath = "/export/data/rsync/fund/app_crm_user_fix_repay_attribute_det.txt.load3"
	//go DealFile(filePath,loadPath)

	//wg.Wait()
}
