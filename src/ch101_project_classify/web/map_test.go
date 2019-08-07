package main

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T){

	names := []string{"Tom", "Jerry"}
	//nums := []string{"one", "two", "three"}
	//pNames := names   // 确认 names 被更新

	var pNames  []string
	//copy(pNames,names)
	copy(names,pNames)


	t.Log(names,pNames)
	//pNames[1] = "test"

	t.Log(names,pNames)
}

func TestMap2(t *testing.T){

	names := []int{1,3}
	//nums := []string{"one", "two", "three"}
	//pNames := names   // 确认 names 被更新

	var pNames  = make ([]int,2)
	copy(pNames,names)

	names[1] = 888
	t.Log(names,pNames)
	pNames[1] = 999

	t.Log(names,pNames)
}

func TestArr1(t *testing.T) {
	var a []int
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	var b = make([]int, 10)
	copy(b, a)
	a[0] = 999
	fmt.Println(b[0])

	fmt.Println(a,b)
}