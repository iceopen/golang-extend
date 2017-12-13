package main

import (
	"log"
	"time"

	"fmt"

	"github.com/nsqio/go-nsq"
)

func main() {

	cfg := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("shtelecom-logs", "infoepoch", cfg)
	if err != nil {
		log.Fatal(err)
	}
	// 设置消息处理函数
	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		fmt.Println(string(message.Body[:]))
		time.Sleep(5 * time.Second)
		return nil
	}))
	// 连接到单例nsqd
	if err := consumer.ConnectToNSQD("172.16.50.143:4150"); err != nil {
		log.Fatal(err)
	}
	//if err := consumer.ConnectToNSQD("127.0.0.1:4150"); err != nil {
	//	log.Fatal(err)
	//}
	<-consumer.StopChan
}
