package main

import (
	"fmt"
	"math"
)

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}
func main(){
	fmt.Println(hypot(3,4)) // "5"

	fmt.Println(add(1,2))
	fmt.Println(add(1,2,3))
	fmt.Println(add([]int{1,2,3}...))
}


func add(args ...int)int{
	var sum int
	for _,v := range args{
		sum += v
	}
	return sum
}