package main

import (
	"fmt"
	"regexp"
)

const text = `my email is 2tsetr@qq.com
my email is lzrumik@qq.com
my email is ccs234@163.com
my email is 2tsetr@qq.com
my email is 2tsetr@qqaacom
my email is mmm@qq.com`

func main(){
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)\.([a-zA-Z0-9]+)`)

	match := re.FindAllStringSubmatch(text,-1)
	fmt.Println(match)
}
