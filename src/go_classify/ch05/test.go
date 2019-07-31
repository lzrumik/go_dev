package main

import "fmt"

func main() {
	fmt.Println(double1(5))
	fmt.Println(double1(6))
	fmt.Println()
	fmt.Println(double2(5))
	fmt.Println(double2(6))
}

// 匿名返回
// 加倍参数，若结果超过 10 则还原
func double1(v1 int) int {
	var v2 int
	v2 = v1 * 2
	defer func() {
		if v2 > 10 {
			v2 = v1 // v2 不会被修改
		}
	}()

	return v2
}

// 有名返回
func double2(v1 int)(v2 int) {
	// v2 与函数一起被声明，在 defer 中能被修改
	defer func() {
		if v2 > 10 {
			v2 = v1 // v2 被修改
		}
	}()

	v2 = v1 * 2
	return
}