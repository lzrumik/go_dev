package main

import (
	"fmt"
	"unicode/utf8"
)

func main(){
	var s string = "yes 我爱慕课网!"

	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	for k,v := range s {
		fmt.Printf("%d %X   ",k,v)
	}

	fmt.Println()
	byteTest := []byte(s)
	for k,v := range byteTest{
		fmt.Println(k,string(v))
	}

	runeTest := []rune(s)
	for k,v := range runeTest{
		fmt.Printf("%d %c ",k,v)
		fmt.Println(k,string(v))
	}
}
