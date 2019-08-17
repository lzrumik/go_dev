package main

import (
	"fmt"
	"math"
	"math/cmplx"
)


/**
z1=(a,b),z2=(c,d)
z1 + z2=(a+c,b+d)
z1 × z2=(ac-bd,bc+ad)
 */
func main(){
	var x complex128 = complex(1, 2) // 1+2i
	var y complex128 = complex(3, 4) // 3+4i
	fmt.Println(x,y)                 // "(-5+10i)"
	fmt.Println(x*y)                 // "(-5+10i)"
	fmt.Println(real(x*y))           // "-5"
	fmt.Println(imag(x*y))           // "10"
}

func main33(){
	fmt.Println(cmplx.Exp(1i*math.Pi)+1)
	fmt.Printf("%.3f\n",cmplx.Exp(1i*math.Pi)+1)

	main2()
}


func main2(){
	var a,b int = 3,4
	var c int = int(math.Sqrt(float64(a*a+b*b)))
	fmt.Println(c)
}