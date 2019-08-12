package main

import "testing"

/**
go test add_test.go add.go -v -cover
 */
func TestAdd(t *testing.T) {
	except := 30
	x,y := 10,20
	sum := Add(x,y)
	if except != sum {
		t.Errorf("Add %d  + %d = %d,except:%d",x,y,sum,except)
	}
}
