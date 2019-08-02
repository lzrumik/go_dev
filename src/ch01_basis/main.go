package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	a := 1
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {         //goroutine1
		defer wg.Done()
		a = a + 1
	}()

	go func() {         //goroutine2
		defer wg.Done()
		if a==1 {
			runtime.Gosched()
			fmt.Println("a==",a)
		}
	}()

	runtime.Gosched()
	wg.Wait()
}