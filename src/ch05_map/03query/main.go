package main

import "fmt"

func main(){
	var ages map[string]int
	fmt.Println(ages == nil)    // "true"
	fmt.Println(len(ages) == 0) // "true"

	//必须得make  否则报错
	ages = make(map[string]int)

	ages["carol"] = 21 // panic: assignment to entry in nil map


	if age, ok := ages["bob"];ok {
		fmt.Println(age)
	}else{
		fmt.Println("bob 不存在")
	}
}
