package main

import (
	"regexp"
	"testing"
)

func OriginPcre(){
	r, _ := regexp.Compile(`([_a-z0-9]+?)([_a-z0-9-]*)@([a-z0-9-]*)(\.[a-z]{2,})`)
	r.FindStringSubmatch("fdsa@2345.com")
}

//直接使用原生的
func BenchmarkOriginPcre(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		OriginPcre()
	}
}

func OriginPcre2(){
	var a,b = 1,2
	c := a +b
	_ = c
}

//直接使用原生的
func BenchmarkAddStringWithOperator(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		OriginPcre2()
	}
}
