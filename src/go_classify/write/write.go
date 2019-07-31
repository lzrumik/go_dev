package main

import (
	"bufio"
	_ "bufio"
	"flag"
	"log"
	"os"
	"time"
)

func main() {
	var filePath string

	flag.StringVar(&filePath, "f", "", "-f=path/to/your/filePath")
	flag.Parse()

	if filePath == "" {
		filePath = os.Getenv("HOME") + "/temp-go.txt"
	}
	filePath = `D:\workspace_go\src\go_classify\write\tmp.txt`

	file, _ := os.Create(filePath)
	defer file.Close()
	os.Truncate(filePath,0)

	wt := bufio.NewWriter(file)

	//写入1百万行数据，11s
	sTime := time.Now().UnixNano() / 1e6
	for i := 0; i < 1e7; i++ {
		_, err := wt.WriteString("helllo worldhelllo worldhelllo worldhelllo worldhelllo worldhelllo worldhelllo worldhelllo world\n")
		//_, err := file.WriteString("helllo worldhelllo worldhelllo worldhelllo worldhelllo worldhelllo worldhelllo worldhelllo world\n")
		if err != nil {
			log.Println(err)
		}
	}

	wt.Flush()
	eTime := time.Now().UnixNano() / 1e6
	log.Printf("cost time:%d/ms", (eTime - sTime))
}