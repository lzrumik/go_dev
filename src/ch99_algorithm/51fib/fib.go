package main

import (
	"fmt"
	"os"
)

/**
递归实现斐波那契
 */
func fib_recursion(num int) int{
	if num == 1 || num == 2{
		return 1
	}

	return fib_recursion(num -1)+fib_recursion(num-2)
}

/**
数组方法  每次计算保存结果
占用一定的空间
 */
func fibarray(term int) []int {
	farr := make([]int, term)
	farr[0], farr[1] = 1, 1

	for i:= 2; i < term; i++ {
		farr[i] = farr[i-1] + farr[i-2]
	}
	return farr
}

/**
闭包方法
相当于 每次记录结果 back1  back2
 */
func fib_closure() func() int {
	back1, back2 := 0, 1
	return func() int {
		// 重新赋值
		back1, back2 = back2, (back1 + back2)
		return back1
	}
}

const LIM = 500

func main(){

	fmt.Println(fibarray(LIM))

	f := fib_closure() //返回一个闭包函数
	var array [LIM]int
	for i := 0; i < LIM; i++ {
		array[i] = f()
	}
	fmt.Println(array)

	os.Exit(0)

	// 递归求第50个数已经  很慢了
	for i:=1;i<=LIM;i++{
		fmt.Printf("fib_recursion %d = %d \n",i,fib_recursion(i))
	}
}
