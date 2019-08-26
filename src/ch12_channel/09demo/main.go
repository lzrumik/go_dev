package main

import (
	"fmt"
	"time"
)


func work(id int ,c chan int){
	//for n := range c{
	//	fmt.Printf("Worker %d received  %c \n",id,n)
	//}

	for {
		n,ok := <- c
		if !ok{
			break
		}
		fmt.Printf("Worker %d received  %c \n",id,n)
	}
}

func createWork(id int , )chan<- int{
	c := make(chan int )
	go work(id,c)
	return c
}


func chanDemo(){
	//var c chan int // c == nil

	var channels [10]chan<- int

	for i:=0;i<10;i++{
		channels[i] = createWork(i)
	}


	for i:=0;i<10;i++{
		channels[i] <- 'a'+ i
	}
	for i:=0;i<10;i++{
		channels[i] <- 'A'+ i
	}

	time.Sleep(time.Millisecond)
}


func bufferedChannel(){
	c := make(chan int,3)

	go work(0,c)

	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	//c <- 4

	time.Sleep(time.Millisecond)
}


func channelClose(){
	c := make(chan int,3)

	go work(0,c)

	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	//c <- 4

	//close(c)
	time.Sleep(time.Millisecond)
}

func main(){
	//chanDemo()
	//
	//bufferedChannel()

	channelClose()
}
