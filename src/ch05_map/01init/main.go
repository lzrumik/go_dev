package main

import "fmt"

func main(){
	var ages map[string]int
	fmt.Println(ages == nil)    // "true"
	fmt.Println(len(ages) == 0) // "true"
}


func main2(){
	ages1 := make(map[string]int) // mapping from strings to ints

	fmt.Println(ages1)

	ages2 := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	fmt.Println(ages2)

	ages3 := make(map[string]int)
	ages3["alice"] = 31
	ages3["charlie"] = 34
	fmt.Println(ages3)

	ages4 := map[string]int{}
	fmt.Println(ages4)

	ages4["alice"] = 32
	fmt.Println(ages4["alice"]) // "32"

	delete(ages4, "alice") // remove element ages["alice"]
	fmt.Println(ages4["alice"]) // "0"
	fmt.Printf("%v\n",ages4["alice"]) // "0"

	ages4["bob"] = ages4["bob"] + 1 // happy birthday!
	ages4["bob"] += 1
	ages4["bob"]++
	fmt.Println(ages4)

	for name, age := range ages4 {
		fmt.Printf("%s\t%d\n", name, age)
	}
}
