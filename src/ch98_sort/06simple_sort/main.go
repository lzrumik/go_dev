package main

import "fmt"

func simple_sort(a []int) {

	for i := 0; i < len(a); i++ {
		var min int = i
		for j := i + 1; j < len(a); j++ {
			if a[min] > a[j] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}
}

/**
简单排序
每趟从当前待排序列中选出最小的放在已拍好序列的末尾
 */
func main() {
	b := [...]int{8, 7, 5, 4, 3, 10, 15}
	simple_sort(b[:])
	fmt.Println(b)
}
