package main

import (
	"runtime"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/parnurzeal/gorequest"
)

func doGet(c *context.Context) {
	url := "http://127.0.0.1:9410" + c.Input.URL()
	bodyBytes := c.Input.RequestBody
	bodyStr := string(bodyBytes)
	logs.Info("请求地址：", c.Input.URL())
	logs.Info("请求地址：", url)
	logs.Info("请求内容：", bodyStr)
	request := gorequest.New()
	resp, body, _ := request.Post(url).
		Send(bodyStr).
		End()
	logs.Info("zipkin请求返回状态：", resp.StatusCode, "内容：", body)
	c.WriteString(body)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	beego.BConfig.RunMode = "prod"
	beego.BConfig.CopyRequestBody = true
	beego.Any("*", doGet)
	beego.Run(":9411")
}
