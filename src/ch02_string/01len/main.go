package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func len_demo(){

	s := "hello, world"
	fmt.Println(len(s))     // "12"
	fmt.Println(s[0], s[7]) // "104 119" ('h' and 'w')

	fmt.Println(s[0:5]) // "hello"


	fmt.Println(s[:5]) // "hello"
	fmt.Println(s[7:]) // "world"
	fmt.Println(s[:])  // "hello, world"

	fmt.Println("goodbye" + s[5:]) // "goodbye, world"
}

func compare(){

	s := "left foot"
	t := s
	s += ", right foot"
	fmt.Println(s) // "left foot, right foot"
	fmt.Println(t) // "left foot"

	s1 := t
	fmt.Println(s1==t) //字符串可以比较
	fmt.Println(s==t) //字符串可以比较

	//s[0] = 'L' // 尝试修改字符串内部数据的操作也是被禁止的 compile error: cannot assign to s[0]
}

func prefix_suffix(){

	s := "left foot"
	s2 := "foot"

	fmt.Println(strings.HasPrefix(s,s2))  //前缀
	fmt.Println(strings.HasSuffix(s,s2))  //后缀
	fmt.Println(strings.Contains(s,s2))   //包含
}

// https://books.studygolang.com/gopl-zh/ch3/ch3-05.html
func utf8_demo(){
	//字符串包含13个字节，以UTF8形式编码，但是只对应9个Unicode字符
	s := "Hello, 世界"
	fmt.Println(len(s))                    // "13"
	fmt.Println(utf8.RuneCountInString(s)) // "9"

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		// https://golang.org/pkg/fmt/  %c unicode 编码
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	for k,v := range s{
		fmt.Println(k,string(v))
	}
	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}


	fmt.Printf("%U\n",'然')  // 7136
	fmt.Printf("%U\n",'京')  // 4EAC

	fmt.Println(string(65))     // "A", not "65"
	fmt.Println(string(0x4eac)) // "京"
	fmt.Printf("%U\n",string(0x4eac)) // "京"
	fmt.Printf(string(0x7136)) // 然
}

func main(){
	//len_demo()
	//compare()
	//prefix_suffix()
	//utf8_demo()
}
