package main

import "fmt"

func main() {
	a := 1
	b := 2
	defer add("A", a, add("B", a, b))
	a = 0
	defer add("C", a, add("D", a, b))
	b = 1
}


func add(desc string, a, b int) int {
	sum := a + b
	fmt.Println(desc, a, b, sum)
	return sum
}
