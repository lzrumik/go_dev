package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func main(){

	var wg sync.WaitGroup
	for i:=0;i<10;i++{
		wg.Add(1)
		go geturl(&wg)
	}

	wg.Wait()
}


func geturl(wg *sync.WaitGroup) string {
	//url := "http://127.0.0.1:8880/user/login"
	url := "http://www.baidu.com"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("get err:", err)
		return ""
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("get data err:", err)
		return ""
	}
	res.Body.Close()

	wg.Done()
	fmt.Println(string(data))
	return string(data)
}