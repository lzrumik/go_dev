package _1panic_recover

import "fmt"

func main() {
	f(3)
}
func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)

	defer func(){
		if p := recover();p!=nil{
			fmt.Println("panic===",p)
			return
		}
	}()

	f(x - 1)
}
