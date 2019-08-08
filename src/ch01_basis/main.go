package main

import (
	"fmt"
	"time"
)

type Line struct {
	Text string
	Time time.Time
	Err  error // Error from tail
}


type Tail struct {
	Lines    chan *Line
}


func main() {


	time := time.Now()
	time.Format("2006-01-02 15:04:05")
	str := `10.254.124.110 - - [01/Apr/2019:16:17:27 +0800] "GET /assets/jquery.min.js HTTP/1.1" 304 0 "http://10.255.72.27:9111/activity/by_date.html" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36"`



	t := &Tail{
		Lines:    make(chan *Line,10),
	}


	line := &Line{}
	line.Text = str
	line.Time = time


	t.Lines <- line

	msg := <- t.Lines

	//fmt.Printf("msg: %s  %s \n",line.Text , line.Time)
	fmt.Printf("msg: %s  %s \n",msg.Text , msg.Time)
}