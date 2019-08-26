package main

import (
	"fmt"
	"time"
)

/**
go run -race main.go
 */
func main(){
	var a [10]int

	// 100个线程  1000个异步io处理
	for i:=0;i<10;i++{
		go func(ii int){ // race condition !
			for {
				fmt.Printf("hello from goroutine %d \n",ii)
				//a[ii]++
				//runtime.Gosched()
			}
		}(i)
	}

	time.Sleep(time.Minute)
	fmt.Println(a)
}
