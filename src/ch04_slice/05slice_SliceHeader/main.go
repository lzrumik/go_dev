package main

import (
	"fmt"
	"reflect"
	"unsafe"
)
//参考 https://mp.weixin.qq.com/s/V4VXrhKzyw9r_1FIzH5sWw
//Go语言slice的本质-SliceHeader

type Slice []int

func (A Slice)Append(value int) {
	fmt.Printf("=========1== %p %d %d \n",A,len(A),cap(A))
	A1 := append(A, value)
	fmt.Printf("=========2== %p %d %d \n",A1,len(A1),cap(A1))


	sh:=(*reflect.SliceHeader)(unsafe.Pointer(&A))
	fmt.Printf("A  Data:%d,Len:%d,Cap:%d\n",sh.Data,sh.Len,sh.Cap)

	sh1:=(*reflect.SliceHeader)(unsafe.Pointer(&A1))
	fmt.Printf("A1 Data:%d,Len:%d,Cap:%d\n",sh1.Data,sh1.Len,sh1.Cap)

}

func main() {
	//mSlice := make(Slice, 10, 20)
	mSlice := make(Slice, 10, 10)
	mSlice.Append(5)
	fmt.Println(mSlice)
}


