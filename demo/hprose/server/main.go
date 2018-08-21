package main

import (
	"github.com/hprose/hprose-golang/rpc"
	"time"
)

func hello(name string) string {
	return "Hello " + name + "!"
}

func nowTime(name string) string {
	return "Hello2 " + name + "!"
}

func tcpServer() {
	server := rpc.NewTCPServer("tcp4://0.0.0.0:4321/")
	server.AddFunction("hello", hello)
	server.Publish("ip", 0, 0)
	server.Timeout = 5 * time.Second
	server.Start()
}

func main() {
	tcpServer()
	//service := rpc.NewHTTPService()
	//service.AddFunction("hello", hello)
	//service.AddFunction("nowTime", nowTime)
	//beego.Handler("/path", service)
	//beego.Run()
}
