package main

import (
	"fmt"
	"reflect"
	"time"
)

//定时器  1s执行一次功能
func main(){
	t := time.NewTicker(time.Second)

	fmt.Printf("%s\n",reflect.TypeOf(t.C))

	for v := range t.C {
		fmt.Println("hello",v.Format("2006-01-02 15:04:05"))
	}
}
