package main

import (
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"github.com/nsqio/go-nsq"
	"log"
	"github.com/astaxie/beego/orm"
	"fmt"
	"time"
	"strconv"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:@tcp(172.16.50.143:4000)/shtelecom?charset=utf8", 30)
	orm.Debug = true

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
		log.Println(string(message.Body))
		i := time.Now().Unix()
		res, err := o.Raw(" INSERT INTO openid_active (openid, acttime, type, note) VALUES(?, ?, ?, ?); ", string(message.Body), strconv.FormatInt(i, 10), "BIND", "WX").Exec()
		if err == nil {
			num, _ := res.RowsAffected()
			fmt.Println("mysql row affected nums: ", num)
		}
		time.Sleep(3 * time.Millisecond)
		return nil
	}))
	// 连接到单例nsqd
	if err := consumer.ConnectToNSQD("172.16.50.143:4150"); err != nil {
		log.Fatal(err)
	}
	<-consumer.StopChan
}
