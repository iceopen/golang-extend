package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"github.com/nsqio/go-nsq"
)

type ActBody struct {
	Openid  string `json:"openid"`  // 主体内容
	Acttime string `json:"acttime"` // 活动时间
	Type    string `json:"type"`    // 类型
	Note    string `json:"note"`    // 备注
	Uuid    string `json:"uuid"`    // UUID
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/shtelecom?charset=utf8", 30)
	orm.Debug = false
}

func main() {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	cfg := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("shtelecom-active", "mohoo", cfg)
	if err != nil {
		log.Fatal(err)
	}
	// 设置消息处理函数
	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		var actBody ActBody
		errJson := json.Unmarshal(message.Body, &actBody)
		if errJson != nil {
			return err
		}
		_, err := o.Raw("INSERT INTO openid_active (openid, acttime, type, note, uuid) VALUES(?, ?, ?, ?, ?); ", actBody.Openid, actBody.Acttime, actBody.Type, actBody.Note, actBody.Uuid).Exec()
		if err == nil {
			log.Println("id:", actBody.Openid, " time:", actBody.Acttime)
		} else {
			log.Fatal(err.Error())
			panic(1)
		}
		time.Sleep(5 * time.Millisecond)
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
