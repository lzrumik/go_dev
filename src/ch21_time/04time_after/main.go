package main

import (
	"fmt"
	"time"
)

func main(){
	//定时一次
	fmt.Println(time.Now())
	t := time.After(time.Second)
	<- t
	fmt.Println(time.Now())
}
