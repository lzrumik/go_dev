package main

import "fmt"

func main(){
	t := [3]int{1,2,3}

	s := t[0:]
	fmt.Println(s)

	t[2] = 10
	fmt.Println(t,s)

	a := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println(reverse(a[:])) //[5 4 3 2 1 0]
}


//int数组翻转
func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}