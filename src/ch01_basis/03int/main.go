package main

import "fmt"

//溢出问题
func int_overflow(){
	var u uint8 = 255
	fmt.Println(u, u+1, u*u) // "255 0 1"

	var i int8 = 127
	fmt.Println(i, i+1, i*i) // "127 -128 1"
}

func main()  {
	int_overflow()


	x := 1<<1 | 1<<5  //按位与  10  10000  =》 10010
	fmt.Println(x)
	fmt.Printf("%08b\n",x)
}
