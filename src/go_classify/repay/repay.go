package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)
 
var wg sync.WaitGroup

func getDate(datetime string) string {
	//fmt.Println(datetime[:10],reflect.TypeOf(datetime[:10]))
	//os.Exit(-1)
	// 2006-01-02 15:04:05.0
	return datetime[:10]
	//datetime := "2015-01-01 00:00:00.0"  //待转化为时间戳的字符串

	timeLayout := "2006-01-02 15:04:05"  //转化所需模板
	loc, _ := time.LoadLocation("Local") //获取时区
	tmp, _ := time.ParseInLocation(timeLayout, datetime, loc)
	timestamp := tmp.Unix() //转化为时间戳 类型是int64

	data := time.Unix(timestamp, 0).Format("2006-01-02")
	//fmt.Println(time.Unix(timestamp, 0).Format("2006-01-02"))
	return data
}

func DealFile(filePath string,loadPath string,typeMap map[int]int){

	loadFileHandle, err := os.Create(loadPath)
	if err != nil {
		panic(err)
	}
	defer loadFileHandle.Close()
	os.Truncate(loadPath, 0)

	//打开文件
	file, _ := os.Open(filePath)
	defer file.Close()

	//创建文件的缓冲读取器
	//reader := bufio.NewReader(file)
	reader := bufio.NewReaderSize(file, 4096 * 1000)
	line := 0
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
		//fund_type := int(repayInfo.CaCd)
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

		if line%3000 == 0 {
			loadFileHandle.WriteString(buffer.String())
			buffer.Reset()
		}
	}

	if buffer.String() != "" {
		loadFileHandle.WriteString(buffer.String())
	}

	fmt.Println("deal file over, deal db")


	fmt.Println("init TRUNCATE ",time.Now().UTC())
	testSql := "TRUNCATE activity_subdivision_search_invest_repay_logger_reality_innodb;"
	fmt.Println(testSql)
	result ,err:= db.Exec(testSql)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected())


	fmt.Println("init LOAD ",time.Now().UTC())
	testSql = "LOAD DATA LOCAL INFILE '"+loadPath+"'  INTO TABLE activity_subdivision_search_invest_repay_logger_reality_innodb (uid,repay_date,repay_amount,type,loan_id);"
	mysql.RegisterLocalFile(loadPath)
	fmt.Println(testSql)
	result ,err = db.Exec(testSql)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected())


	wg.Done()
}

var db *sql.DB

func init() {
	fmt.Println(time.Now().UTC())
	db, _ = sql.Open("mysql", "beta-mysql:nopass.2@tcp(10.255.72.27:3306)/crm_activity?charset=utf8")
	db.SetMaxOpenConns(20) //用于设置最大打开的连接数，默认值为0表示不限制。
	db.SetMaxIdleConns(20)//用于设置闲置的连接数。设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	err := db.Ping()
	if err != nil {
		fmt.Println(err)
		panic("connect mysql error")
	}

	fmt.Println("init db ",time.Now().UTC())

	testSql := "insert into activity_subdivision_search_invest_repay_logger_reality_innodb values(0,123,\"2019-07-06\",1.23,1,\"hahshfsdfjdsaljfalsdf \")"
	result ,err:= db.Exec(testSql)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected())


	rows,err := db.Query("select * from activity_subdivision_search_invest_repay_logger_reality_innodb order by id desc limit 3")
	if err != nil {
		panic(err)
	}
	for rows.Next(){
		var id int
		var uid int
		var repay_date string
		var repay_amount float64
		var repay_type int
		var loan_id string
		err = rows.Scan(&id,&uid, &repay_date, &repay_amount, &repay_type, &loan_id)

		fmt.Println(id,uid,repay_date,repay_amount,repay_type,loan_id)
	}
	rows.Close()
}


func main() {

	//创建输出文件
	//f, err := os.Create("cpu.prof")
	//if err != nil {
	//	log.Fatal("could not create CPU profile:", err)
	//}
	//
	//if err := pprof.StartCPUProfile(f); err != nil {
	//	log.Fatal("could not start cpu profile", err)
	//}
	//
	//defer pprof.StopCPUProfile()


	filePath := "/export/data/rsync/fund/app_crm_user_fix_repay_attribute_det.txt"
	loadPath := "/export/data/rsync/fund/app_crm_user_fix_repay_attribute_det.txt.load"

	//filePath = `D:\workspace_go\src\go_classify\repay\a.txt`
	//loadPath = `D:\workspace_go\src\go_classify\repay\a.load`

	fmt.Println("init dealfile ",time.Now().UTC())
	typeMap := map[int]int{1:1,2:2}
	wg.Add(1)
	go DealFile(filePath,loadPath,typeMap) 


	//filePath = "/export/data/rsync/fund/app_crm_user_p2p_repay_attribute_det.txt"
	//loadPath = "/export/data/rsync/fund/app_crm_user_fix_repay_attribute_det.txt.load2"
	//typeMap = map[int]int{1:3,2:4}
	//wg.Add(1)
	//go DealFile(filePath,loadPath,typeMap)

	//f1, err := os.Create("mem.prof")
	//if err != nil {
	//	log.Fatal("could not create mem profile:", err)
	//}
	//
	//if err := pprof.WriteHeapProfile(f1); err != nil {
	//	log.Fatal("could not start mem profile", err)
	//}
	//f1.Close()
	//
	//f2, err := os.Create("goroutine.prof")
	//if err != nil {
	//	log.Fatal("could not create goroutine profile:", err)
	//}
	//
	//if gProf := pprof.Lookup("goroutine"); gProf == nil {
	//	log.Fatal("could not start goroutine profile", err)
	//} else {
	//	gProf.WriteTo(f2, 0)
	//}
	//f2.Close()

	wg.Wait()

	//filePath = "/export/data/rsync/fund/app_crm_user_p2p_repay_attribute_det.txt"
	//loadPath = "/export/data/rsync/fund/app_crm_user_fix_repay_attribute_det.txt.load2"
	//typeMap = map[int]int{1:3,2:4}
	//go DealFile(filePath,loadPath,typeMap)
	//
	//
	//filePath = "/export/data/rsync/fund/app_crm_fengchu_repay_det.txt"
	//loadPath = "/export/data/rsync/fund/app_crm_user_fix_repay_attribute_det.txt.load3"
	//go DealFile(filePath,loadPath)

}
