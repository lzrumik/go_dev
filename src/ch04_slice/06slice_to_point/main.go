package main

import (
	"fmt"
	"unsafe"
)

//从slice中得到一块内存地址
func getSlicePoint(){
	s := make([]byte, 200)
	ptr := unsafe.Pointer(&s[0])

	fmt.Println(ptr)
}

func main(){
	getSlicePoint()

	fmt.Println(1<<10)

}
