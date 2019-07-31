package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1) // 强制使多个 goroutine 串行执行
	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println("i: ", i)
			wg.Done()
		}()
		// time.Sleep(1 * time.Second)  // 此时将顺序输出 1 2 3 4 5
	}

	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
