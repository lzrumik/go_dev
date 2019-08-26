package main

import (
	"errors"
	"fmt"
	"os"
)

type queuearr struct {
	maxSize int
	data    [5]int
	front   int
	rear    int
}

func (t *queuearr) addQueue(val int) (flag bool, err error) {

	if t.rear == t.maxSize {
		return false, errors.New("queue is full")
	}
	t.data[t.rear] = val
	t.rear++

	return true, nil
}

func (t *queuearr) getQueue() (val int, err error) {
	if t.front == t.rear {
		return -1, errors.New("queue is empty")
	}

	//fmt.Println(t.data)
	val = t.data[t.front]
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
			fmt.Println(val)
			_, err := testS.addQueue(val)
			dealErr(err)
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
