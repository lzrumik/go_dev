package main

import (
	"fmt"
	"unsafe"
)


// https://books.studygolang.com/gopl-zh/ch13/ch13-01.html
func main(){
	fmt.Println(unsafe.Sizeof(float64(0))) // "8"
}
