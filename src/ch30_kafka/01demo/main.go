package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Shopify/sarama"
)


/**

 */

func main() {

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err := sarama.NewSyncProducer([]string{"192.168.7.19:9092"}, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		return
	}

	defer client.Close()
	for i:=0;i<1000;i++ {
		msg := &sarama.ProducerMessage{}
		msg.Topic = "nginx_log"
		msg.Value = sarama.StringEncoder("this is a good test, my message is good"+strconv.Itoa(i))

		pid, offset, err := client.SendMessage(msg)
		if err != nil {
			fmt.Println("send message failed,", err)
			return
		}

		fmt.Printf("pid:%v offset:%v\n", pid, offset)
		time.Sleep(10 * time.Millisecond)
	}
}
