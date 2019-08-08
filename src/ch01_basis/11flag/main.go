package main

import (
	"flag"
	"fmt"
	"os"
)

/**
go build .
11flag.exe -c=1 -d=2
 */
func main() {
	fmt.Printf("len of args:%d\n", len(os.Args))
	for i, v := range os.Args {
		fmt.Printf("args[%d]=%s\n", i, v)
	}

	var confPath string
	var logLevel int
	flag.StringVar(&confPath, "c", "", "please input conf path")
	flag.IntVar(&logLevel, "d", 10, "please input log level")

	flag.Parse()

	fmt.Println("path:", confPath)
	fmt.Println("log level:", logLevel)
}
