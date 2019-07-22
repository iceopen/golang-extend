package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nsqio/go-nsq"
)

var producer *nsq.Producer

// 主函数
func main() {
	fmt.Printf("时间戳（秒）：%v;\n", time.Now().Unix())
	fmt.Println("============")
	fmt.Printf("时间戳（纳秒）：%v;\n", time.Now().UnixNano())
	fmt.Println("============")

	startTime := time.Now().UnixNano()

	Publish("speed", "speed")

	Consumer()
	endTime := time.Now().UnixNano()
	fmt.Println("============")
	fmt.Println(endTime - startTime)
}

// InitProducer 初始化生产者
func InitProducer(str string) {
	var err error
	fmt.Println("address: ", str)
	producer, err = nsq.NewProducer(str, nsq.NewConfig())
	if err != nil {
		panic(err)
	}
}

// Publish 发布消息
func Publish(topic string, message string) error {
	strIP1 := "127.0.0.1:4150"
	InitProducer(strIP1)
	var err error
	if producer != nil {
		if message == "" { //不能发布空串，否则会导致error
			return nil
		}
		i := 1
		for {
			err = producer.Publish(topic, []byte(message+string(i))) // 发布消息
			i++
			if i > 10*10000 {
				break
			}
		}
		return err
	}
	return fmt.Errorf("producer is nil", err)
}

// Consumer 消费速度测试
func Consumer() {
	cfg := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("speed", "speed-1", cfg)
	if err != nil {
		log.Fatal(err)
	}
	// 设置消息处理函数
	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		fmt.Println(" -- " + string(message.Body[:]))
		return nil
	}))
	// 连接到单例nsqd
	if err := consumer.ConnectToNSQD("127.0.0.1:4150"); err != nil {
		log.Fatal(err)
	}
	<-consumer.StopChan
}
