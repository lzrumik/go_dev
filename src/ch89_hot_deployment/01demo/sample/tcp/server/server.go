package main

import (
	"ch89_hot_deployment/01demo/reload"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"sync/atomic"
	"syscall"
)

var (
	port     int
	stopChan = make(chan struct{})
)

/**
参考 https://github.com/514366607/reload
kill -SIGTSTP pid    关闭
kill -SIGHUP pid     重启

kill -SIGHUP  `pidof ./server` #重启

https://github.com/fvbock/endless  2296 star
https://github.com/facebookarchive/grace 4373 star
https://github.com/jpillora/overseer 13101

 */
func main() {
	flag.IntVar(&port, "p", 8888, `端口`)
	flag.Parse()

	log.Printf("Actual pid is %d\n", syscall.Getpid())

	listener, err := reload.GetListener(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Println(err)
	}

	var s = reload.NewService(listener)
	log.Printf("isChild : %v ,listener: %v\n", s.IsChild(), listener)

	go func() {
		defer listener.Close()

		var isAccept int32 = 1

		go func() {
			select {
			case <-stopChan:
				atomic.StoreInt32(&isAccept, 2)
			}
		}()

		for atomic.LoadInt32(&isAccept) == 1 {
			conn, err := listener.Accept()
			if err != nil {
				log.Println(err)
				continue
			}
			s.Add(1)
			log.Println("Accept ", conn.RemoteAddr())
			go recvConnMsg(conn, s)
		}
	}()

	s.Start()

	stopChan <- struct{}{}
}

func recvConnMsg(conn net.Conn, s reload.Service) {
	//  var buf [4096]byte
	buf := make([]byte, 4096)

	defer conn.Close()
	defer s.Done()

	for {
		n, err := conn.Read(buf)

		if err == io.EOF {
			//连接结束
			return
		} else if err != nil {
			log.Println(err)
			return
		}

		var recv = string(buf[0:n])
		recv = fmt.Sprintf(" pid %d Return: %s ", syscall.Getpid(), recv)
		log.Printf("Rev Data : %v", recv)

		conn.Write([]byte(recv))

	}
}