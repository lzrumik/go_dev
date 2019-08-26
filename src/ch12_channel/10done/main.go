package main

import (
	"fmt"
)


func work(id int ,c chan int,done chan bool){
	for n := range c{
		fmt.Printf("Worker %d received  %c \n",id,n)
		go func(){
			done <- true
		}()
	}

	//for {
	//	n,ok := <- c
	//	if !ok{
	//		break
	//	}
	//	fmt.Printf("Worker %d received  %c \n",id,n)
	//}
}

type worker struct {
	in chan int
	done chan bool
}

func createWork(id int , )worker{
	w := worker{
		in : make( chan int),
		done : make(chan bool),
	}

	go work(id,w.in,w.done)
	return w
}


func chanDemo(){
	//var c chan int // c == nil

	var workers [10]worker

	for i:=0;i<10;i++{
		workers[i] = createWork(i)
	}


	for i,worker := range workers{
		worker.in <- 'a'+ i
	}
	for _,worker := range workers{
		<-worker.done
	}
	for i,worker := range workers{
		worker.in <- 'A'+ i
	}

	for _,worker := range workers{
		<-worker.done
	}
}

func main(){
	chanDemo()
	//
	//bufferedChannel()

	//channelClose()
}
