package main

import "fmt"

func fab(num int) int{
	if num == 1 || num == 2{
		return 1
	}

	return fab(num -1)+fab(num-2)
}


func main(){
	fmt.Println(fab(1))
	fmt.Println(fab(2))
	fmt.Println(fab(3))
	fmt.Println(fab(4))
	fmt.Println(fab(5))
	fmt.Println(fab(6))
	fmt.Println(fab(7))
	fmt.Println(fab(16))
	fmt.Println(fab(46))



}
