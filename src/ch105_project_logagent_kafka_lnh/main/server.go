package main

import (
	"ch105_project_logagent_kafka_lnh/kafka"
	"ch105_project_logagent_kafka_lnh/tailf"
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
	//"fmt"
)

func serverRun() (err error) {
	for {
		msg := tailf.GetOneLine()
		err = sendToKafka(msg)
		if err != nil {
			logs.Error("send to kafka failed, err:%v", err)
			time.Sleep(time.Second)
			continue
		}
	}
	return
}

func sendToKafka(msg *tailf.TextMsg) (err error) {
	fmt.Printf("read msg:%s, topic:%s\n", msg.Msg, msg.Topic)
	err = kafka.SendToKafka(msg.Msg, msg.Topic)
	return
}
