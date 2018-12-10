package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
)

func OpenTrance(c *context.Context) {
	bodyBytes := c.Input.RequestBody
	bodyStr := string(bodyBytes)

	logs.Info("请求地址：", c.Input.URI())
	logs.Info("请求内容：", bodyStr)
	c.WriteString("")
}

// 启动方法
func main() {
	beego.Any("*", OpenTrance)
	beego.Run(":80")
}
