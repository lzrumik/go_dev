package main

import (
	"fmt"
)

func main(){
	var a1 [9]int
	fmt.Println(a1)//默认0值   [0 0 0 0 0 0 0 0 0]

	//三种定义方式
	var a = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var b [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//数组越界检查
	var a2 = [5]int{1,2,3,4,5}
	//a2[101] = 255// invalid array index 101 (out of bounds for 5-element array)
	fmt.Println(a2)

	//数组赋值 赋值的两个数组变量的值不会共享
	var a3 = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var b3 [9]int
	b3 = a3
	a3[0] = 12345
	fmt.Println(a3)
	fmt.Println(b3)

	//数组遍历
	var a4 = [5]int{1,2,3,4,5}
	for index := range a4 {
		fmt.Println(index, a4[index])
	}
	for index, value := range a4 {
		fmt.Println(index, value)
	}
}
