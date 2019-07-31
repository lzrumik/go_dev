package main

import (
	"fmt"
	"runtime"
)

func  main42()  {

	runtime.GOMAXPROCS(2)
	for i:=0;i<1000;i++ {
		go fmt.Print(0)
		fmt.Print(1)
	}
}


func main(){
	p1 := new(int)
	fmt.Printf("p1 --> %#v \n ", p1) //(*int)(0xc42000e250)
	fmt.Printf("p1 point to --> %#v \n ", *p1) //0

}