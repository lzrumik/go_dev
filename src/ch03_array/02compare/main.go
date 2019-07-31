package main

import (
	"crypto/sha256"
	"fmt"
)

func sha256_test(){
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)  // %t bool值
}

func main(){
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c) // "true false false"
	d := [3]int{1, 2}
	fmt.Println(d)
	//fmt.Println(a == d) // compile error: cannot compare [2]int == [3]int

	sha256_test()
}
