package main

import (
	"fmt"
	"sync"
)


func work(id int ,c chan int,done func()){
	for n := range c{
		fmt.Printf("Worker %d received  %c \n",id,n)
		done()
	}
}

type worker struct {
	in chan int
	done func()
}

func createWork(id int ,wg_demo *sync.WaitGroup )worker{
	w := worker{
		in : make( chan int),
		done : func(){
			wg_demo.Done()
		},
	}

	go work(id,w.in,w.done)
	return w
}


func chanDemo(){
	//var c chan int // c == nil

	var workers [10]worker

	var wg sync.WaitGroup
	wg.Add(20)

	for i:=0;i<10;i++{
		workers[i] = createWork(i,&wg)
	}



	for i,worker := range workers{
		worker.in <- 'a'+ i
	}
	for i,worker := range workers{
		worker.in <- 'A'+ i
	}

	wg.Wait()

}

func main(){
	chanDemo()
	//
	//bufferedChannel()

	//channelClose()
}
