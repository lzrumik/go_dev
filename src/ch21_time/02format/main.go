package main

import (
	"fmt"
	"time"
)

func main(){
	loc,_ := time.LoadLocation("Asia/Shanghai")
	timeStr := "29/Mar/2019:10:39:34 +0800"
	fmt.Printf("%q\n",timeStr)

	t ,err := time.ParseInLocation("02/Jan/2006:15:04:05 -0700",timeStr,loc)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(t.Format("2006-01-02 15:04:05"))
}
