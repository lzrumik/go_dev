package _2table_test

import (
	"math"
	"testing"
)

func calcTriangle(a,b int)int{
	return int(math.Sqrt(float64(a*a+b*b)))
}

/*
表格驱动测试
go test main_test.go
go test .
 */
func TestTriangle(t *testing.T){
	tests := []struct{a,b,c int}{
		{3,4,5},
		{5,12,0},
		{8,15,17},
		{12,35,0},
		{30000,40000,50000},
	}

	for _,tt := range tests{
		if actual := calcTriangle(tt.a,tt.b);actual != tt.c{
			t.Errorf("calcTriangle（%d %d) got %d ; excepted %d",tt.a,tt.b,tt.c,actual)
		}
	}
}
