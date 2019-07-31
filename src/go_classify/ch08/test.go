package main

import (
	"fmt"
	"reflect"
)
type People struct {
	Name string
}

func (p *People) String() string {
	return fmt.Sprintf("print: %v", p)
}

func main() {
	p := People{}
	//p :=& People{}

	fmt.Println(reflect.TypeOf(p))
	fmt.Printf("%p",p)
	//p.String()
}