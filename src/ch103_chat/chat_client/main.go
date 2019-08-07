package main

import (
	"fmt"
	"ch103_chat/proto"
	"net"
)

var userId int
var passwd string
var msgChan chan proto.UserRecvMessageReq

func init() {
	msgChan = make(chan proto.UserRecvMessageReq, 1000)
}

func main() {
	//var userId int
	//var passwd string

	conn, err := net.Dial("tcp", "localhost:10020")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}

	fmt.Println("please input the username password")
	fmt.Scanf("%d %s\n", &userId, &passwd)

	err = login(conn, userId, passwd)
	if err != nil {
		fmt.Println("login failed, err:", err)
		return
	}

	go processServerMessage(conn)
	for {
		logic(conn)
	}
}
