package main

import (
	"fmt"
	"unsafe"
)

func sum(a *int){
	fmt.Println(unsafe.Pointer(&a))
}

func sum2(a int){
	fmt.Println(unsafe.Pointer(&a))
}
func main(){
	var a int = 5
	fmt.Println(unsafe.Pointer(&a))
	sum(&a)
	sum2(a)
}
