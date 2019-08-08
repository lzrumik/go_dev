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

type Student struct {
	Name string
	Sex  string
}

func Test(a interface{}) {
	b, ok := a.(Student)
	if ok == false {
		fmt.Println("convert failed")
		return
	}
	//b += 3
	fmt.Println(b)
}

func just(items ...interface{}) {
	for index, v := range items {
		switch v.(type) {
		case bool:
			fmt.Printf("%d params is bool, value is %v\n", index, v)
		case int, int64, int32:
			fmt.Printf("%d params is int, value is %v\n", index, v)
		case float32, float64:
			fmt.Printf("%d params is float, value is %v\n", index, v)
		case string:
			fmt.Printf("%d params is string, value is %v\n", index, v)
		case Student:
			fmt.Printf("%d params student, value is %v\n", index, v)
		case *Student:
			fmt.Printf("%d params *student, value is %v\n", index, v)
		}
	}
}


func main(){
	fmt.Println(grade(100))
	fmt.Println(grade(85))
	fmt.Println(grade(75))
	fmt.Println(grade(55))
	//fmt.Println(grade(-1))


	var b Student = Student{
		Name: "stu01",
		Sex:  "female",
	}
	Test(b)
	just(28, 8.2, "this is a test", b, &b)
}
