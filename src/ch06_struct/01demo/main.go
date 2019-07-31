package main

import (
	"fmt"
)
type People struct {
	Name string
}

func (p People) String() string {
	p.Name = "55555"
	return p.Name
	//return fmt.Sprintf("%v",p)
}
func (p *People) String2() string {
	p.Name = "6666"
	return p.Name
	//return fmt.Sprintf("%v",p)
}
//传值与传指针
func main() {
	p := People{Name:"12334"}
	fmt.Println(p.String())   //传值
	fmt.Println(p.Name)

	fmt.Println(p.String2())  //传指针
	fmt.Println(p.Name)
}