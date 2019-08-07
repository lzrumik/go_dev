package main

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

type queuearr struct {
	maxSize int
	data    [5]int
	front   int
	rear    int
}

func (t *queuearr) addQueue(val int) (flag bool, err error) {

	fullFlag := false
	for _,v := range t.data{
		if v!=0{
			fullFlag = true
			break
		}
	}
	if fullFlag{
		fmt.Println(t.data,t.front,t.rear)
		return false, errors.New("queue is full")
	}
	t.data[t.rear] = val
	t.rear++
	t.rear = t.rear%t.maxSize

	return true, nil
}

func (t *queuearr) getQueue() (val int, err error) {
	flag := false
	for _,v := range t.data{
		if v!=0{
			flag = true
			break
		}
	}

	if !flag {
		fmt.Println(t.data,t.front,t.rear)
		return -1, errors.New("queue is empty")
	}

	val = t.data[t.front]
	t.data[t.front] = 0
	t.front++

	return val, nil
}

func (t *queuearr) showQueue() (val []int, err error) {
	if t.front == t.rear {
		return val, errors.New("queue is empty")
	}

	return t.data[t.front:], nil
}

func main() {
	var testS = queuearr{
		maxSize: 5,
	}
	var key string
	var val int

	for {
		fmt.Println("1,输入add 表示添加数据到队列")
		fmt.Println("2,输入get 表示从队列获取数据")
		fmt.Println("3,输入show 表示从队列获取所有数据")
		fmt.Println("4,输入exit 表示退出")
		fmt.Scanf("%s", &key)

		switch key {
		case "add":
			fmt.Println("请输入要插入的数据")
			fmt.Scanf("%d", &val)
			if val<=0{
				fmt.Println("请输入大于0的整数")
				break;
			}
			fmt.Println("您输入的数字是:",val)
			_, err := testS.addQueue(val)
			dealErr(err)

			//fmt.Println(testS)
		case "show":
			fmt.Println(testS.showQueue())
		case "get":
			val, err := testS.getQueue()
			if err != nil {
				fmt.Println(err)
				break
			}

			fmt.Printf("从队列取出的数据是=%d \n", val)
		case "exit":
			os.Exit(0)
		}
	}
}

func dealErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
