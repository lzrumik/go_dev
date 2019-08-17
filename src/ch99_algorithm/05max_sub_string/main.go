package main

import "fmt"


/**
获取一段字符串中 最长的子串

定义一个map 记录上次出现字母的index
定义一个start 记录子串的第一个字母的index2
 */
func getMaxSubString(str string)(int,string){
	lastCurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i,ch := range []rune(str){
		ch_last_index,ok := lastCurred[ch]
		//如果 该字母在子串的右边出现 更新start
		if ok && ch_last_index >= start {
			start = lastCurred[ch] + 1
		}

		//如果start发生变化   就更新maxlength
		if i-start +1 > maxLength {
			maxLength = i - start +1
		}

		maxLength = i - start +1

		lastCurred[ch] = i
	}

	subString := ([]rune(str))[start:(start+maxLength)]
	return maxLength,string(subString)
}

func main(){

	fmt.Println(getMaxSubString("hahalira2"))
	fmt.Println(getMaxSubString("bbbbbbb"))
	fmt.Println(getMaxSubString("我爱慕课爱一231 "))
	fmt.Println(getMaxSubString("一二三二一"))

}
