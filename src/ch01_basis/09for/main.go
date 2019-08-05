package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string){
	file ,err := os.Open(filename)
	if err!=nil{
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println(convertToBin(5))      //101
	fmt.Println(convertToBin(13))     //1101
	fmt.Println(convertToBin(111111)) //1101

	printFile("D:/github/go_dev/src/ch01_basis/09for/abc.txt")

	content, err := ioutil.ReadFile("D:/github/go_dev/src/ch01_basis/09for/abc.txt")

	if err!=nil{
		panic(err)
	}
	fmt.Printf("%s \n",content)
}
