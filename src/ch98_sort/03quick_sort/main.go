package main

import (
	"fmt"
)

func QuickSort(values []int) {
	if len(values) <= 1 {
		return
	}
	mid, i := values[0], 1  //获取第一个数为 基准值
	head, tail := 0, len(values)-1  //最开始的key  和 最末尾的key
	for head < tail {
		fmt.Println(values)
		if values[i] > mid {  //从第二个元素开始比较基准值 如果大于基准值 就交换元素至末尾
			values[i], values[tail] = values[tail], values[i]
			tail--
		} else { // 如果小于基准值 就交换元素至开头
			values[i], values[head] = values[head], values[i]
			head++
			i++
		}
	}
	values[head] = mid
	QuickSort(values[:head])
	QuickSort(values[head+1:])
}

/**
快速排序：先选定一个基准元素，然后以该基准元素划分数组，再在被划分的部分重复以上过程，最后可以得到排序结果。
[8 4 5 7 1 3 6 2]
[8 4 5 7 1 3 6 2]
[4 8 5 7 1 3 6 2]
[4 5 8 7 1 3 6 2]
[4 5 7 8 1 3 6 2]
[4 5 7 1 8 3 6 2]
[4 5 7 1 3 8 6 2]
[4 5 7 1 3 6 8 2]
[4 5 7 1 3 6 2]
[4 2 7 1 3 6 5]
[2 4 7 1 3 6 5]
[2 4 6 1 3 7 5]
[2 4 3 1 6 7 5]
[2 3 4 1 6 7 5]
[2 3 1]
[2 1 3]
[6 7 5]
[6 5 7]
[1 2 3 4 5 6 7 8]
 */
func main(){
	//array := []int{55, 94, 87, 1, 4, 32, 11, 77, 39, 42, 64, 53, 70, 12, 9}
	array := []int{8,4,5,7,1,3,6,2}
	fmt.Println(array)
	QuickSort(array)
	fmt.Println(array)
}
