package main

import (
	"fmt"
	"reflect"
)

func main(){
	var slice []int  //nil slice  声明了slice，却没有给实例化的对象
	slice2 := make([]int,0)// empty slice

	fmt.Println(slice,slice2)

	fmt.Println(reflect.DeepEqual(slice, slice2))
}
