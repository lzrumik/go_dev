package main

import "fmt"

func test1() {
	// make 函数创建切片，需要提供三个参数，分别是切片的类型、切片的长度和容量。
	var s1 []int = make([]int, 5, 8)
	var s2 []int = make([]int, 8) // 满容切片
	fmt.Println(s1)
	fmt.Println(s2)
}

func test2(){
	var s1 []int //nil 切片
	var s2 []int = []int{}  // 空切片
	var s3 []int = make([]int, 0)
	fmt.Println(s1, s2, s3)
	fmt.Println(len(s1), len(s2), len(s3))
	fmt.Println(cap(s1), cap(s2), cap(s3))
}


func test4(){
	var s1 = []int{1,2,3,4,5}
	fmt.Println(s1, len(s1), cap(s1))

	// 对满容的切片进行追加会分离底层数组
	var s2 = append(s1, 6)
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))

	// 对非满容的切片进行追加会共享底层数组
	var s3 = append(s2, 7)
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println(s3, len(s3), cap(s3))
}

func test5(){
	s1 := make([]int, 6)
	s2 := make([]int, 1024)
	s1 = append(s1, 1)
	s2 = append(s2, 2)
	fmt.Println(len(s1), cap(s1))
	fmt.Println(len(s2), cap(s2))
}

func main(){
	test5()
}
