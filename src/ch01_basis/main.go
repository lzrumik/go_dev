package main

import "fmt"

type Point struct { X, Y int }

type Rect1 struct { Min, Max Point }
type Rect2 struct { Min, Max *Point }

func main() {

	p := Point{10,20}
	pp := &Point{10,20}

	fmt.Println(pp,p)
	fmt.Printf("%p %d\n",pp,p)

	r1 := Rect1{Point{10,20},Point{50,60}}
	r2 := Rect2{&Point{10,20},&Point{50,60}}

	fmt.Println(r1,r2)


	fmt.Println("=================================")

	tSlice := []int{1,2,3,4,5}
	fmt.Println(tSlice)





	var r *Rect = &Rect{10,20}
	fmt.Println(r)
}


type Rect struct {
	width float64
	height float64
}

