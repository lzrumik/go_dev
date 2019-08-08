package main

import (
	"errors"
	"fmt"
	"github.com/labstack/gommon/log"
)

func add(x,y int) (r int,err error){
	if x<0 || y <0 {
		return  -1,fmt.Errorf("param %d < 0", x)
	}

	return x+y,nil
}

func main(){
	sum ,err := add(-1,1)
	if err!=nil{
		log.SetPrefix("wait:")
		//log.SetHeader("header")
		log.Fatalf("%s",errors.New("add error : "+err.Error()))
	}

	fmt.Println("sum==",sum)
}