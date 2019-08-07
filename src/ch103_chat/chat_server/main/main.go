package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("server start!")
	initRedis("localhost:6379", 16, 1024, time.Second*300)

	pool.Get().Do("hset","users","1",`{"user_id":1,"passwd":"123456789","nick":"join2","sex":"female","header":"leader","last_login":"2017-01-02","status":1}`)

	fmt.Println("update done!")
	initUserMgr()
	runServer("0.0.0.0:10020")
}
