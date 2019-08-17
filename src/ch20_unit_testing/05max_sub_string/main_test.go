package main

import (
	"testing"
)

/**
覆盖率
 */
func TestSubstr(t *testing.T){
	tests := []struct{
		s string
		ans int
	}{
		//Normal case
		{"abcabcbb",3},
		{"pwwkew",3},


		//Edge cases
		{"abcabcbb",3},
		{"pwwkew",3},

		//chinese support
		{"我爱慕课爱一231 ",8},
		{"一二三二一",3},
	}



	for _,tt := range tests{

		actual,_,str := getMaxSubString(tt.s)
		if actual != tt.ans{
			t.Errorf("getMaxSubString（%s) got %d ; excepted %d, %s",tt.s,actual,tt.ans,str)
		}
	}
}


/**
go test -bench=.
go test -bench . -cpuprofile cpu.out
web

*/
func BenchmarkSubstr(b *testing.B){
	s := "黑化肥挥发发灰"
	for i:=0;i<13;i++{
		s = s + s
	}

	b.ResetTimer()

	ans := 6
	for i := 0;i<b.N;i++{
		actual,_,_ := getMaxSubString(s)
		if actual != ans {
			b.Errorf("got %d for input %s; expected %d",actual,s ,ans)
		}
	}
}

func BenchmarkSubstr2(b *testing.B){
	s := "黑化肥挥发发灰"
	for i:=0;i<13;i++{
		s = s + s
	}
	b.ResetTimer()

	ans := 6
	for i := 0;i<b.N;i++{
		actual,_,_ := getMaxSubStringNew(s)
		if actual != ans {
			b.Errorf("got %d for input %s; expected %d",actual,s ,ans)
		}
	}
}