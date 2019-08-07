package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

var c redis.Conn
/**
docker run --name myredis -d -p 6379:6379 redis
 */
func main(){

	var err error
	c, err = redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}

	defer c.Close()

	fmt.Println("conn redis success!")

	set()
	get()
	expire()

	time.Sleep(time.Second*5)
	get()
}


func set(){
	_, err := c.Do("Set", "abc", 100)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func get(){
	r, err := redis.Int(c.Do("Get", "abc"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}

	fmt.Println(r)
}


func expire(){
	_, err := c.Do("expire", "abc", 3)
	if err != nil {
		fmt.Println(err)
		return
	}
}