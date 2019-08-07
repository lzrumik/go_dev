package main

import (
	"fmt"
	"os"
	"text/template"
)

type Person struct {
	Name  string
	Title string
	age   string
}

func main() {
	t, err := template.ParseFiles("index.html")
	//t, err := template.ParseFiles("/Users/kevint/github/go_dev/src/ch22_template_text/02template/index.html")
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	p := Person{Name: "Mary", age: "31", Title: "我的个人网站"}
	if err := t.Execute(os.Stdout, p); err != nil {
		fmt.Println("There was an error:", err.Error())
	}
}
