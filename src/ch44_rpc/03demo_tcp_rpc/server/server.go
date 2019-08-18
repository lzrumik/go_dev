package main

import (
	"errors"
	"log"
	"net"
	"net/rpc"
)

//声明矩形对象
type Rect struct {

}

//声明参数结构体
type Params struct {
	Width,Height int
}


//求矩形面积
func (r *Rect) Area(p Params,ret *int) error{
	*ret = p.Width * p.Height
	return nil
}

//求矩形周长
func(r *Rect)Perimeter(p Params,ret *int) error{
	*ret = (p.Width + p.Height) * 2
	return nil
}



//声明Arith对象
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

func (p *ArithRequest) Multiply(req ArithRequest,res *ArithResponse)error{
	res.Pro = req.A * req.B
	return nil
}

func (p *ArithRequest) Divide(req ArithRequest,res *ArithResponse)error{
	if req.B == 0 {
		return errors.New("除数不能为0")
	}
	res.Que = req.A / req.B
	res.Rem = req.A % req.B
	return nil
}


//RPC服务端
func main(){
	//1、注册服务
	rect := new(Rect)
	rpc.Register(rect)

	rpc.Register(new(ArithRequest))

	tcpAddr ,err := net.ResolveTCPAddr("tcp",":18080")
	if err!= nil {
		log.Fatal(err)
	}
	listener,err := net.ListenTCP("tcp",tcpAddr)
	if err!= nil {
		log.Fatal(err)
	}

	for{
		conn,err := listener.Accept()
		if err!= nil {
			log.Fatal("conn err :",err)
		}
		rpc.ServeConn(conn)
	}


	////2、把服务绑定到http协议上
	//rpc.HandleHTTP()
	////3、监听服务
	//err := http.ListenAndServe(":8080", nil)
	//if err!= nil {
	//	log.Fatal(err)
	//}
}
