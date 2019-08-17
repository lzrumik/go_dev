package main

import "fmt"


/**
获取一段字符串中 最长的子串

定义一个map 记录上次出现字母的index
定义一个start 记录子串的第一个字母的index2
 */
func getMaxSubString(str string)(int,int,string){
	lastCurred := make(map[rune]int)
	start := 0
	wordStart := 0
	maxLength := 0

	for i,ch := range []rune(str){

		//如果 该字母在子串的右边出现 更新start
		if lastI,ok := lastCurred[ch];ok && lastI >= start {
			start = lastI + 1
		}

		//如果start发生变化   就更新maxlength
		if i-start+1 > maxLength {
			wordStart = start
			maxLength = i - start +1
		}

		lastCurred[ch] = i
	}
	subString := ([]rune(str))[wordStart:(wordStart+maxLength)]
	return maxLength,wordStart,string(subString)
}


func main(){

	//                              黑化肥发灰会挥
	//s := "化肥会挥发黑化肥发灰灰化肥发黑黑化肥发灰会挥发"
	s := "黑化肥挥发发灰"
	for i:=0;i<3;i++{
		s = s + s
	}
	fmt.Println(getMaxSubString(s))
	fmt.Println(getMaxSubStringNew(s))


	//fmt.Println(getMaxSubString("hahalira2"))
	//fmt.Println(getMaxSubString("bbbbbbb"))
	//fmt.Println(getMaxSubString("我爱慕课爱一231 "))
	//fmt.Println(getMaxSubString("一二三二一"))

}

var lastCurred = make([]int ,0xffff)
func getMaxSubStringNew(str string)(int,int,string){
	// stores last occurred pos + 1
	start := 0
	maxLength := 0
	wordStart := 0
	for i := range lastCurred{
		lastCurred[i] = 0
	}

	for i,ch := range []rune(str){

		//如果 该字母在子串的右边出现 更新start
		if  lastI := lastCurred[ch] ; lastI > start {
			start = lastI
		}

		//如果start发生变化   就更新maxlength
		if i-start +1 > maxLength {
			wordStart = start
			maxLength = i - start
		}

		lastCurred[ch] = i
	}


	//fmt.Println([]rune(str),len([]rune(str)))
	subString := ([]rune(str))[wordStart:(wordStart+maxLength)]
	return maxLength,wordStart,string(subString)
}
