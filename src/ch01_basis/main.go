package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//测试字符串连接
func main2(){
	str := "hello"
	var s string
	start := time.Now()
	for i := 0 ; i < 100000; i++ {
		s += str
	}
	end := time.Since(start)

	fmt.Println(end)

	var s2 strings.Builder
	start2 := time.Now()
	for i := 0 ; i < 100000; i++ {
		s2.WriteString(str)
	}
	s2.String()
	end2 := time.Since(start2)

	fmt.Println(end2)
}


//测试字符串连接
func main3(){
	var s strings.Builder
	//var s string;
	start := time.Now()
	for i := 0 ; i < 100000; i++ {

		//s += strconv.Itoa(i)
		s.WriteString(strconv.Itoa(i))
		//fmt.Println(s)
	}
	end := time.Since(start)

	fmt.Println(end)
}


//测试字符串连接
func main(){
	var s2 strings.Builder
	start2 := time.Now()
	for i := 0 ; i < 1000000; i++ {
		s2.WriteString(strconv.Itoa(i))
	}
	s2.String()
	end2 := time.Since(start2)

	fmt.Println(end2)

	var s3 bytes.Buffer
	start3 := time.Now()
	for i := 0 ; i < 1000000; i++ {
		s3.WriteString(strconv.Itoa(i))
	}
	s3.String()
	end3 := time.Since(start3)

	fmt.Println(end3)


	//var s string
	//start := time.Now()
	//for i := 0 ; i < 1000000; i++ {
	//	s += strconv.Itoa(i)
	//}
	//end := time.Since(start)
	//
	//fmt.Println(end)

}
