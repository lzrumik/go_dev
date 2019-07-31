package main

import (
	"fmt"
	"os"
)

// 获取命令行参数 go run main.go 1 2 3
func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}