package main

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
)

func getUrlBody() chan string {
	str := make(chan string)
	fmt.Println("url start")
	go func() {
		fmt.Println("httplib start")
		urlStr, err := httplib.Get("http://httpbin.org/get").String()
		if err != nil {

		}
		fmt.Println("httplib end")
		str <- urlStr
	}()
	fmt.Println("url end")
	return str
}

func strPrint(input chan string) {
	for ss := range input {
		fmt.Println(ss)
	}
}

func main() {
	str := getUrlBody()
	fmt.Println(str)
	go strPrint(str)
	select {}
}
