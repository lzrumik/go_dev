package _9func

import "fmt"

func main042(){
	a := func(){
		fmt.Println("learn go")
	}
	a()

	fmt.Printf("%T\n",a)
	fmt.Println("over")

	i := 100
	func(){
		i = 1000
		fmt.Println("the i value is",i)
	}()
	fmt.Println("the i value is",i)
}


type add func(a int,b int)int

func main(){
	var a add = func(a int,b int)int{
		return a + b
	}
	s := a(5,6)
	fmt.Println("Sum", s)
}