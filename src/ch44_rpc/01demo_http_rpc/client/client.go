package main

import (
	"fmt"
	"log"
	"net/rpc"
)

//声明矩形对象
type Rect struct {

}

//声明参数结构体
type Params struct {
	Width,Height int
}

type ArithRequest struct {
	A,B int
}

type ArithResponse struct {
	//乘积
	Pro int
	//商
	Que int
	//余数
	Rem int
}

//RPC客户端调用服务
func main(){
	//1、远程连接RPC服务
	conn,err := rpc.DialHTTP("tcp",":18081")
	if err != nil{
		log.Fatal(err)
	}

	//求面积
	//定义一个接受服务端传回来的计算结果的变量
	ret := 0
	p := Params{Width:10,Height:20}
	conn.Call("Rect.Area",p,&ret)
	fmt.Println("面积:",ret)


	//求周长
	//定义一个接受服务端传回来的计算结果的变量
	ret2 := 0
	conn.Call("Rect.Perimeter",p,&ret2)
	fmt.Println("周长:",ret2)


}
