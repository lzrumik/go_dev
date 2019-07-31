package main

import "fmt"

func main(){
	s := [...]int{1,2,3,4,5}

	s1 := s[3:]
	s2 := s[3:]
	fmt.Println(s1,s2)
	//fmt.Println(s1==s2) // slice 不能比较

	var t []int
	if t == nil { //true
		fmt.Println("slice is nil")
	}

	var t1 = []int{}
	if t1 == nil { //true
		fmt.Println("slice is nil")
	}
}
