package main

import "fmt"

func main(){
	type Point struct{ X, Y int }

	p := Point{1, 2}
	q := Point{2, 1}
	fmt.Println(p.X == q.X && p.Y == q.Y) // "false"
	fmt.Println(p == q)                   // "false"


	q = Point{1, 2}
	fmt.Println(p.X == q.X && p.Y == q.Y) // "true"
	fmt.Println(p == q)                   // "true"
}
