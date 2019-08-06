package main

import "fmt"

func MergeSort(array []int) []int{
	n := len(array)
	if n < 2 {
		return array
	}
	key := n / 2
	left := MergeSort(array[0:key])
	right := MergeSort(array[key:])
	return merge(left, right)
}

func merge(left []int, right []int) []int{
	newArr := make([]int, len(left)+len(right))
	i, j, index :=0,0,0
	for {
		if left[i] > right[j] {
			newArr[index] = right[j]
			index++
			j++
			if j == len(right) {
				copy(newArr[index:], left[i:])
				break
			}

		}else{
			newArr[index] = left[i]
			index++
			i++
			if i == len(left) {
				copy(newArr[index:], right[j:])
				break
			}
		}
	}
	return newArr
}

/**
归并排序:先将数组不断细分成最小的单位，然后每个单位分别排序，排序完毕后合并，重复以上过程最后就可以得到排序结果。

https://www.cnblogs.com/chengxiao/p/6194356.html  算法复杂度 O(nlog n)

https://zh.wikipedia.org/wiki/%E5%BD%92%E5%B9%B6%E6%8E%92%E5%BA%8F
 */
func main() {
	//array := []int{55, 94, 87, 1, 4, 32, 11, 77, 39, 42, 64, 53, 70, 12, 9}
	array := []int{8,4,5,7,1,3,6,2}
	fmt.Println(array)
	array = MergeSort(array)
	fmt.Println(array)

}