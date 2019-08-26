package main

import "fmt"

func main() {
outer:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == 5 {
				break outer
			}
			fmt.Printf("i,j= %d %d \t", i, j)
		}
		fmt.Println()
	}
}
