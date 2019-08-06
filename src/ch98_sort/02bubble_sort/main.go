package main

import (
	"fmt"
)

func BubbleSort(values []int) {
	flag := true
	vLen := len(values)
	for i := 0; i < vLen-1; i++ {
		flag = true
		for j := 0; j < vLen-i-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				flag = false
				continue
			}
		}
		if flag {
			break
		}
	}
}

/*
	冒泡排序 O(n*n) n的平方
	依次比较两个相邻的元素,比较大小后如果不是顺序就交换位置，直到没有相邻元素需要交换
	冒泡排序是一种稳定排序算法

[8 4 5 7 1 3 6 2]
[4 8 5 7 1 3 6 2]
[4 5 8 7 1 3 6 2]
[4 5 7 8 1 3 6 2]
[4 5 7 1 8 3 6 2]
[4 5 7 1 3 8 6 2]
[4 5 7 1 3 6 8 2]
[4 5 7 1 3 6 2 8]
[4 5 7 1 3 6 2 8]
[4 5 7 1 3 6 2 8]
[4 5 1 7 3 6 2 8]
[4 5 1 3 7 6 2 8]
[4 5 1 3 6 7 2 8]
[4 5 1 3 6 2 7 8]
[4 5 1 3 6 2 7 8]
[4 1 5 3 6 2 7 8]
[4 1 3 5 6 2 7 8]
[4 1 3 5 6 2 7 8]
[4 1 3 5 2 6 7 8]
[1 4 3 5 2 6 7 8]
[1 3 4 5 2 6 7 8]
[1 3 4 5 2 6 7 8]
[1 3 4 2 5 6 7 8]
[1 3 4 2 5 6 7 8]
[1 3 4 2 5 6 7 8]
[1 3 2 4 5 6 7 8]
[1 3 2 4 5 6 7 8]
[1 2 3 4 5 6 7 8]
[1 2 3 4 5 6 7 8]
 */
func main(){
	//array := []int{55, 94, 87, 1, 4, 32, 11, 77, 39, 42, 64, 53, 70, 12, 9}
	array := []int{8,4,5,7,1,3,6,2}
	fmt.Println(array)
	BubbleSort(array)
	fmt.Println(array)
}
