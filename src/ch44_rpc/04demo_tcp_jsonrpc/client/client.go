package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

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
	conn,err := jsonrpc.Dial("tcp",":8081")
	if err != nil{
		log.Fatal(err)
	}

	req := ArithRequest{100,0}
	var res ArithResponse
	err2 := conn.Call("ArithRequest.Multiply",req,&res)
	if err2!= nil{
		log.Fatal(err2)
	}
	fmt.Printf("%d * %d  = %d \n",req.A,req.B,res.Pro)

	err3 := conn.Call("ArithRequest.Divide",req,&res)
	if err3 != nil {
		log.Fatal(err3.Error())
	}

	fmt.Printf("%d / %d ，商 = %d ,余数 = %d的\n",req.A,req.B,res.Que,res.Rem)
}
