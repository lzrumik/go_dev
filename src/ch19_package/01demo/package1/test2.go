package package1

import "fmt"

var (
	test = 1
)

func init (){
	fmt.Println("package1 init ",test)
	test = 2
	fmt.Println(test)
}


func Add(){
	fmt.Println(1+2)
}