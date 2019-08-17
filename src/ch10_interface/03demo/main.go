package main

import (
	"ch10_interface/03demo/mock"
	"ch10_interface/03demo/real"
	"fmt"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string{
	return r.Get("http://www.imooc.com")
}

func main(){
	var r Retriever
	r = mock.Retriever{"this is a fack imooc.com"}
	inspect(r)
	//fmt.Printf("%T %v \n",r,r)

	r = &real.Retriever{UserAgent:"Mozilla/5.0",TimeOut:time.Second*2}

	inspect(r)
	//r = &real.Retriever{}
	//fmt.Printf("%T %v \n",r,r)
	//fmt.Println(download(r))


	/**
	  判断对象是不是某个结构体
	*/
	if mockRetriever,ok := r.(mock.Retriever);ok{
		fmt.Println(mockRetriever.Contents)
	}else{
		fmt.Println("is not a mock Retriever")
	}

	fmt.Println(download(mock.Retriever{"this is a fack imooc.com2"}))
}

/**
检测接口的类型
 */
func inspect(r Retriever){
	fmt.Printf("%T %v \n",r,r)

	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("mock Retriever:",v.Contents)
	case *real.Retriever:
		fmt.Println("mock Retriever:",v.UserAgent)
	}
}