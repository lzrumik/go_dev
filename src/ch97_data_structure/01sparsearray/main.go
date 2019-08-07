package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type saveChess struct {
	rowid int
	colid int
	val   int
}

func main() {
	var chessMap [10][10]int
	chessMap[1][2] = 1 //黑子
	chessMap[2][3] = 2 //白子

	chessNum := 1
	for i := 0; i < len(chessMap); i++ {
		for j := 0; j < len(chessMap[i]); j++ {
			fmt.Printf("%d\t", chessMap[i][j])
			if chessMap[i][j] != 0 {
				chessNum++
			}
		}
		fmt.Println()
	}

	fmt.Printf("\n\n\n稀疏数组如下:\n")
	var testArr []saveChess
	//testArr = make([]saveChess, chessNum)
	//testArr = make([]saveChess, 0)
	//testArr[0] = saveChess{0, 0, 0}
	testArr = append(testArr, saveChess{len(chessMap), len(chessMap[0]), 0})

	//var point = 1

	for i := 0; i < len(chessMap); i++ {
		for j := 0; j < len(chessMap[i]); j++ {
			if chessMap[i][j] != 0 {

				testArr = append(testArr, saveChess{i, j, chessMap[i][j]})

				//testArr[point] = saveChess{i, j, chessMap[i][j]}
				//point++
			}
		}
	}

	for i := 0; i < len(testArr); i++ {
		fmt.Println(testArr[i].rowid, testArr[i].colid, testArr[i].val)
	}

	filePath := "/Users/kevint/Downloads/chessmap.data"

	inputFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	inputWriter := bufio.NewWriter(inputFile)
	for i := 0; i < len(testArr); i++ {

		str := strings.Builder{}
		str.WriteString(strconv.Itoa(testArr[i].rowid))
		str.WriteString("\t")
		str.WriteString(strconv.Itoa(testArr[i].colid))
		str.WriteString("\t")
		str.WriteString(strconv.Itoa(testArr[i].val))
		str.WriteString("\n")

		inputWriter.WriteString(str.String())
	}

	if err != nil {
		panic(err)
	}
	inputWriter.Flush()
	fmt.Println("写完毕了")

	var originArr [][]int
	inputFile, err = os.OpenFile(filePath, os.O_RDWR, 0777)
	inputReader := bufio.NewReader(inputFile)
	for {
		teststr, err := inputReader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("读取已结束")
			break
		}
		fmt.Println(teststr)
		teststr = strings.TrimRight(teststr, "\n")
		strArr := strings.Split(teststr, "\t")
		rowId, err := strconv.Atoi(strArr[0])
		if err != nil {
			panic(err)
		}
		colId, err := strconv.Atoi(strArr[1])
		if err != nil {
			panic(err)
		}

		val, err := strconv.Atoi(strArr[2])
		if err != nil {
			panic(err)
		}

		fmt.Println(rowId, colId, val)
		fmt.Println(reflect.TypeOf(colId))
		if val == 0 {
			for i := 0; i < rowId; i++ {
				s1 := make([]int, 0, colId)
				//var s1 []int
				for j := 0; j < colId; j++ {
					s1 = append(s1, val)
				}
				originArr = append(originArr, s1)
			}
			continue
		}

		originArr[rowId][colId] = val
	}

	fmt.Printf("\n\n\n循环输出切片\n")
	for i := 0; i < len(originArr); i++ {
		for j := 0; j < len(originArr[i]); j++ {
			fmt.Printf("%d\t", originArr[i][j])
		}
		fmt.Println()
	}

}
