package main

import "fmt"

func grade(score int) (g string){
	switch  {
	case score < 0 || score>100:
		panic(fmt.Sprintf("Wrong score:%d",score))
	case score<60:
		g = "F"
	case score<80:
		g = "C"
	case score<90:
		g = "B"
	case score<=100:
		g = "A"
	}
	return g
}

func main(){
	fmt.Println(grade(100))
	fmt.Println(grade(85))
	fmt.Println(grade(75))
	fmt.Println(grade(55))
	fmt.Println(grade(-1))
}
