package main

import "fmt"

func main(){
	var numbers [4][5]int
	fmt.Println(numbers)
	for i:=0;i<len(numbers);i++{
		for j:=0;j<len(numbers[i]);j++{
			fmt.Print(numbers[i][j])
		}
		fmt.Println()
	}
}
