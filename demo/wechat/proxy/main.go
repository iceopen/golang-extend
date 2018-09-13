package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"github.com/astaxie/beego/logs"
)

var plainContentType = []string{"text/plain; charset=utf-8"}

func init() {
	logs.SetLogger(logs.AdapterFile, `{"filename":"logs/replay.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	logs.Async()
}

// 返回内容
func writeContextType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}

// 获取发送内容
func getBody(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	con, _ := ioutil.ReadAll(r.Body) // 获取post的数据
	reqUrl := r.RequestURI
	bodyStr := string(con) // 发送的XML内容
	logs.Info(r.RequestURI, "接收到内容：", bodyStr)
	// 请求转发
	contentReader := bytes.NewReader(con)
	reqResByte := make(chan []byte, 1)
	startTime := time.Now().UnixNano()
	go func() {
		// 处理请求
		req, _ := http.NewRequest("POST", "http://172.16.50.131:3040"+reqUrl, contentReader)
		req.Header.Set("Content-Type", "application/xml")
		client := &http.Client{}
		client.Timeout = 10 * time.Second
		resp, _ := client.Do(req)
		defer resp.Body.Close()
		// 获取返回内容返回
		reqRes, _ := ioutil.ReadAll(resp.Body) // 获取post的数据
		logs.Info("返回到内容：", string(reqRes))
		logs.Info("处理耗时：", time.Now().UnixNano()-startTime)
		reqResByte <- reqRes
	}()
	// 返回
	writeContextType(w, plainContentType)
	w.Write(<-reqResByte)
}

func main() {
	// 正式上线地址
	http.HandleFunc("/wxReply", getBody)
	fmt.Println("soft start")
	http.ListenAndServe(":8400", nil)
}
