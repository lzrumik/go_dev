package main

import "fmt"

func main(){
	if (!true==false)==true {
		fmt.Println("true")
	}
	//终于明白 if为啥可以加() 了  因为运算的时候需要指定顺序

	i := 0
	fmt.Println(itob(i))

	j := 1
	fmt.Println(itob(j))
}


func itob(i int) bool { return i != 0 }