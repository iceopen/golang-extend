package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Word!") // 这个写入到w的是输出到客户端的
}

func main() {
	http.HandleFunc("/", sayhelloName)       // 设置访问的路由
	err := http.ListenAndServe(":8080", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	fmt.Println("start ok!")
}
