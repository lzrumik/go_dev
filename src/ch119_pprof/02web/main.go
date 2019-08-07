package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

// 参考 https://studygolang.com/articles/14519?fr=sidebar
// http://127.0.0.1:10002/debug/pprof/
/**
安装Graphviz 配置bin目录到环境变量  重启cmder

go tool pprof http://localhost:10002/debug/pprof/profile

go tool pprof C:\Users\lzrumik\pprof\pprof.samples.cpu.001.pb.gz

 */
func main(){
		log.Println(http.ListenAndServe("localhost:10002", nil))

}