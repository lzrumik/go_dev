package package2

import (
	"ch19_package/01demo/package1"
	"fmt"
)

var (
	test2 = 3
)

func init (){
	fmt.Println("package2 init ",test2)
	test2 = 4
	fmt.Println(test2)
}


func Add2(){
	package1.Add()

	fmt.Println(3+4)
}