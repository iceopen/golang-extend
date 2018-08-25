package main

import (
	"github.com/hprose/hprose-golang/rpc"
	"time"
	"fmt"
)

type HelloService struct {
	Hello   func(string) string
	NowTime func(string) string
}

func httpTest() {
	client := rpc.NewHTTPClient("http://127.0.0.1:8080/path")
	var helloService *HelloService
	client.UseService(&helloService)
	ii := 0
	startTime := time.Now().Unix()
	for {
		helloService.Hello("world1")
		helloService.NowTime("world2")
		ii++
		if ii > 1000 {
			break
		}
	}
	endXX := time.Now().Unix() - startTime
	fmt.Println(endXX)
}

func main() {

}
