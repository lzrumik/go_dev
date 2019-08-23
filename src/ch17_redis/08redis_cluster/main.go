package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

/**
虚拟机开启redis cluster
网络设置端口映射
 */
func main()  {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
	})
	defer client.Close()

	fmt.Println("redis dbsize:",client.DBSize())
	client.Ping()

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	fmt.Println("pool state init state:", client.PoolStats())
	for i := 0; i < 1000; i++ {
		k := fmt.Sprintf("key:%d", i)
		v := k
		val, err := client.Set(k, v, 60*time.Second).Result()
		if err != nil {
			panic(err)
		}

		val, err = client.Get(k).Result()
		if err != nil {
			panic(err)
		}
		fmt.Println("key:", val)
	}
	fmt.Println("pool state final state:", client.PoolStats()) //获取客户端连接池相关信息

}
