package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"time"
)

type Reader interface {
	Read(rc chan []byte)
}
type Writer interface {
	Write(wc chan string)
}

type LogProcess struct{
	rc chan []byte
	wc chan string
	read Reader
	write Writer
}

type ReadFromFile struct {
	path string//读取文件的路径
}

type WriteToInfluxDB struct {
	influxDBDsn string//influx data source
}

func(r *ReadFromFile) Read(rc chan []byte){
	//读取模块
	//打开文件
	f,err := os.Open(r.path)
	if err != nil{
		panic(fmt.Sprintln("open file error:%s",err.Error()))
	}

	//从文件末尾读取数据
	//f.Seek(0,2)
	rd := bufio.NewReader(f)

	for{
		line,err := rd.ReadBytes('\n')
		if err == io.EOF{
			time.Sleep(500 * time.Microsecond)
			break
		}else if err != nil{
			panic(fmt.Sprintln("read file error:%s",err.Error()))
		}
		rc <- line
	}

}

func(w *WriteToInfluxDB) Write(wc chan string){
	//写入模块
	for v := range wc {
		fmt.Println(string(v))
	}
}

func(l *LogProcess) Process(){
	//解析模块
	for v := range l.rc {
		line := `10.254.21.11 - - [29/Mar/2019:10:39:34 +0800] "GET /1.html HTTP/1.1" 200 6 "-" "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36"`

		line = string(v)
		re2str := `^(?P<remote_addr>\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}) .* `+
			`\[(?P<time_local>.*?)\] `+
			`"(?P<request>.*?)" `+
			`(?P<status>[^ ]*) `+
			`(?P<body_bytes_sent>[^ ]*) `+
			`"(?P<http_referer>[^"]*)" ` +
			`"(?P<http_user_agent>[^"]*)"`

		re2 := regexp.MustCompile(re2str)
		testStr := re2.FindStringSubmatch(line)

		regMap := make(map[string]string)
		for k,v := range re2.SubexpNames(){
			key := v
			value := testStr[k]
			regMap[key] = value
		}

		getTime := formatTime(regMap["time_local"])
		getIp   := regMap["remote_addr"]
		request   := regMap["request"]

		l.wc <- getIp + "     " + getTime + "     " + request
	}


	data := <- l.rc
	str := string(data)
	l.wc <- strings.ToUpper(str)
}

func formatTime(timeStr string)string{
	loc,_ := time.LoadLocation("Asia/Shanghai")
	//fmt.Printf("%q\n",timeStr)

	t ,err := time.ParseInLocation("02/Jan/2006:15:04:05 -0700",timeStr,loc)
	if err != nil {
		fmt.Println(err.Error())
	}
	return t.Format("2006-01-02 15:04:05")
}


func main(){

	r := &ReadFromFile{
		path:"access.log",
	}

	w := &WriteToInfluxDB{
		influxDBDsn:"username&password..",
	}

	lp := &LogProcess{
		rc:make(chan []byte),
		wc:make(chan string),
		read:r,
		write:w,
	}

	go lp.read.Read(lp.rc)
	go lp.Process()
	go lp.write.Write(lp.wc)

	time.Sleep(time.Second*5)
}
