package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"strings"
	"time"
)
func main() {
	filename := "access.log"
	//filename := ".\\my.log"
	tails, err := tail.TailFile(filename, tail.Config{
		ReOpen:    true,
		Follow:    true,
		//Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})
	if err != nil {
		fmt.Println("tail file err:", err)
		return
	}
	var msg *tail.Line

	msg, _ = <-tails.Lines

	var ok bool
	for {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		//应该是有个换行的符号
		str := strings.Replace(msg.Text,"\r","\n",-1)
		fmt.Printf("msg: %s  %s \n", str, "hahahhfa")
	}
}
