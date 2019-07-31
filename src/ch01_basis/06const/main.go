package main

import (
	"fmt"
	"math"
)

const (
	Sunday int = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

const (
	e  = 2.71828182845904523536028747135266249775724709369995957496696763
	pi = 3.14159265358979323846264338327950288419716939937510582097494459
)

func main()  {
	const pi = 3.14159 // approximately; math.Pi is a better approximation

	fmt.Println(pi)

	fmt.Println(Saturday)


	const (
		FlagUp int = 1 << iota // is up
		FlagBroadcast            // supports broadcast access capability
		FlagLoopback             // is a loopback interface
		FlagPointToPoint         // belongs to a point-to-point link
		FlagMulticast            // supports multicast access capability
	)

	fmt.Println(FlagMulticast)

	const (
		_ = 1 << (10 * iota)
		KiB // 1024
		MiB // 1048576
		GiB // 1073741824
		TiB // 1099511627776             (exceeds 1 << 32)
		PiB // 1125899906842624
		EiB // 1152921504606846976
		ZiB // 1180591620717411303424    (exceeds 1 << 64)
		YiB // 1208925819614629174706176
	)

	fmt.Println(PiB)

	var x float32 = math.Pi
	var y float64 = math.Pi
	var z complex128 = math.Pi

	fmt.Println(x,y,z)
}
