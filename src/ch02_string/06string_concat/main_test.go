package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)


var (
	str = "hello"
	loop = 5000  //循环拼接才次数
)
func AddStringWithOperator() string{
	strResult := ""
	for j:=0;j<loop;j++{
		strResult += str
	}
	return strResult
}
func AddStringWithSprintf() string{
	strResult := ""
	for j:=0;j<loop;j++{
		strResult = fmt.Sprintf("%s%s", strResult, str)
	}
	return strResult
}

func AddStringWithJoin() string{
	var strSlice []string
	strSlice = make([]string,loop)
	for j:=0;j<loop;j++{
		strSlice[j] = str
	}
	return strings.Join(strSlice, "")
}

func AddStringWithBuffer()string{

	var buffer bytes.Buffer
	for j:=0;j<loop;j++{
		buffer.WriteString(str)
	}
	return  buffer.String()
}

func AddStringWithBuild()string{
	var strb strings.Builder
	for j:=0;j<loop;j++{
		strb.WriteString(str)
	}
	return  strb.String()
}

//func main(){
//
//	fmt.Println(AddStringWithOperator())
//	fmt.Println(AddStringWithSprintf())
//	fmt.Println(AddStringWithJoin())
//	fmt.Println(AddStringWithBuffer())
//	fmt.Println(AddStringWithBuild())
//}
//直接使用运算符
func BenchmarkAddStringWithOperator(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		AddStringWithOperator()
	}
}

//sprintf
func BenchmarkAddStringWithSprintf(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		AddStringWithSprintf()
	}
}

//string join
func BenchmarkAddStringWithJoin(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		AddStringWithJoin()
	}
}


//bytes buffer
func BenchmarkAddStringWithBuffer(b *testing.B) {
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		AddStringWithBuffer()
	}
}


//strings builder
func BenchmarkAddStringWithBuild(b *testing.B){
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		AddStringWithBuild()
	}
}

/**
loop = 100
goos: windows
goarch: amd64
pkg: ch01_basis/06string_concat
BenchmarkAddStringWithOperator-8   	  200000	      9230 ns/op
BenchmarkAddStringWithSprintf-8    	   50000	     22540 ns/op
BenchmarkAddStringWithJoin-8       	 1000000	      1539 ns/op
BenchmarkAddStringWithBuffer-8     	 1000000	      1406 ns/op
BenchmarkAddStringWithBuild-8      	 2000000	       891 ns/op
PASS
 */

/**
loop = 5000
goos: windows
goarch: amd64
pkg: ch01_basis/06string_concat
BenchmarkAddStringWithOperator-8   	     200	   9914993 ns/op
BenchmarkAddStringWithSprintf-8    	     100	  12579974 ns/op
BenchmarkAddStringWithJoin-8       	   20000	     73650 ns/op
BenchmarkAddStringWithBuffer-8     	   30000	     56532 ns/op
BenchmarkAddStringWithBuild-8      	   30000	     42333 ns/op
PASS
*/

/**
结论
性能要求不太高的场合，直接使用运算符，代码更简短清晰，能获得比较好的可读性
直接使用运算符，代码更简短清晰，使用byte.buffer能获得比较好的型男
大量拼接字符串的话，使用 strings.Join byte.buffer能有比较好的性能 byte.buffer好于strings.Join
1.11版本以后 string.builder比byte.buffer能有2倍的性能提升
 */