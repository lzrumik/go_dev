package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch := make(chan struct{})
	go func() {
		fmt.Println("start working")
		time.Sleep(time.Second * 10)
		ch <- struct{}{}
	}()
	go func() {
		fmt.Println("start working2")
	}()

	go func() {
		fmt.Println("start working2")
		time.Sleep(time.Second * 3)
	}()
	go func(){
		for{
			fmt.Println(runtime.NumGoroutine())
			time.Sleep(time.Second)
		}
	}()


	<-ch

	fmt.Println("finished")
}

func main42(){
	ch := make(chan int,2)

	go func(){
		fmt.Println(<-ch)
	}()
	ch<-1
	ch<-2
	ch<-3

	fmt.Println("1111")
}

