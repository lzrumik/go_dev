package main

import (
	"net/http"
	"runtime/pprof"
)

//使用pprof 查看goroutine的数量
//http://localhost:11181/
//参考 ：https://blog.csdn.net/siddontang/article/details/24296943
var quit chan struct{} = make(chan struct{})

func f() {
	<-quit
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	p := pprof.Lookup("goroutine")
	p.WriteTo(w, 1)
}

func main() {
	for i := 0; i < 10000; i++ {
		go f()
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":11181", nil)
}