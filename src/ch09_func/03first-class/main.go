package main

import (
	"fmt"
	"log"
	"strings"
)

func main42(){
	defer func(){
		if p:= recover();p!=nil{
			log.Fatal(p)
		}
	}()
	var f func(int) int
	if f != nil {
		f(3)
	}
	f(3) // 此处f的值为nil, 会引起panic错误
}

func main(){
	//main42()

	fmt.Println(strings.Map(add1, "HAL-9000")) // "IBM.:111"
	fmt.Println(strings.Map(add1, "VMS"))      // "WNT"
	fmt.Println(strings.Map(add1, "Admix"))    // "Benjy"
}


func add1(r rune) rune { return r + 1 }
