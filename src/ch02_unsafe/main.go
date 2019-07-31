package main

import (
	"fmt"
	"reflect"
	"unsafe"
)


type type1 struct{ bool; float64; int16 } // 3 words 4words
type type2 struct{ float64; int16; bool } // 2 words 3words
type type3 struct{ bool; int16; float64 } // 2 words 3words

type type10 struct{ type3; bool; int16; float64 } // 2 words 3words
type type11 struct{ a []int; bool; int16; float64 } // 2 words 3words
func main(){

	fmt.Println(unsafe.Sizeof(int8(0))) // "1"
	fmt.Println(unsafe.Sizeof(int16(0))) // "2"
	fmt.Println(unsafe.Sizeof(int32(0))) // "4"
	fmt.Println(unsafe.Sizeof(int64(0))) // "8"
	fmt.Println(unsafe.Sizeof(int(0))) // "8"
	fmt.Println(unsafe.Sizeof(float64(0))) // "8"
	fmt.Println(unsafe.Sizeof(float32(0))) // "4"


	a := 0
	fmt.Println(unsafe.Sizeof(&a)) // "8"

	str := "1234"
	fmt.Println(reflect.TypeOf(str)) // "8"
	fmt.Println(unsafe.Sizeof(str)) // "8"
	fmt.Println(len(str)) // "8"


	fmt.Println("123---------------------")

	demo1 := type1{true,1,2}
	fmt.Println(unsafe.Sizeof(demo1)) // "1"

	demo2 := type2{1,2,true}
	fmt.Println(unsafe.Sizeof(demo2)) // "1"

	demo3 := type3{true,1,2}
	fmt.Println(unsafe.Sizeof(demo3)) // "1"


	demo10 := type10{demo3,true,1,2}
	fmt.Println(unsafe.Sizeof(demo10.type3)) // "1"
	fmt.Println(unsafe.Sizeof(demo10)) // "1"


	m := make([]int,2)
	m[0] = 10
	m[0] = 10
	demo11 := type11{m,true,1,2}
	fmt.Println(unsafe.Sizeof(m)) // "1"
	fmt.Println(unsafe.Sizeof(demo11)) // "1"
}
