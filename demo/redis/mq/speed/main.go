package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

// 主函数
func main() {
	startTime := time.Now().Unix()
	PushSpeed()
	endTime := time.Now().Unix()
	fmt.Println("============")
	fmt.Printf("耗时（秒）：%v;\n", (endTime - startTime))
}
func newPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
}

var (
	pool        *redis.Pool
	redisServer = flag.String("redisServer", ":6379", "")
)

// PushSpeed 队列推送
func PushSpeed() {
	flag.Parse()
	pool = newPool(*redisServer)
	c := pool.Get()
	i := 1
	for {
		c.Do("PUBLISH", "redisChat", "HELLO")
		i++
		if i > 1*1 {
			break
		}
	}
}

// SubSpeed 消费队列
func SubSpeed() {
	flag.Parse()
	pool = newPool(*redisServer)
	c := pool.Get()
	psc := redis.PubSubConn{Conn: c}
	psc.Subscribe("redisChat")
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			fmt.Println(v)
		}
	}
}
