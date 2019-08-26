package main

import "fmt"

func getNum()interface{}{
	return 1
}

//接口的类型断言
func main(){
	a := getNum()

	var b int = 10
	b = a.(int)

	fmt.Println(a,b)
}
