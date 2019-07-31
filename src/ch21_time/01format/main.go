package main

import (
	"fmt"
	"time"
)

func main(){

	timeStamp := time.Now().Unix()
	fmt.Println("当前时间戳：",timeStamp)

	nowTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("当前时间：",nowTime)

	//时间戳 to 时间
	timeStamp = 140000000
	tm := time.Unix(timeStamp, 0)
	fmt.Println(tm.Format("2006-01-02 15:04:05")) //2018-07-11 15:10:19

	//时间 to 时间戳
	var thisdate = "2019-01-01 12:10:20"
	loc, _ := time.LoadLocation("Asia/Shanghai")        //设置时区
	tt, _ := time.ParseInLocation("2006-01-02 15:04:05", thisdate, loc) //2006-01-02 15:04:05是转换的格式如php的"Y-m-d H:i:s"
	fmt.Println(tt.Unix())
	// http://tool.chinaz.com/Tools/unixtime.aspx

	//格式化时间
	thisdate2 := "2014-03-17 14:55:06.12331"
	tt2, _ := time.ParseInLocation("2006-01-02 15:04:05.9", thisdate2, loc) //2006-01-02 15:04:05是转换的格式如php的"Y-m-d H:i:s"
	fmt.Println(tt2.Format("2006-01-02 15:04:05"))

	fmt.Println(time.Now().Format(time.RFC3339Nano))
}



