package main

import (
	"github.com/elastic/apm-agent-go/module/apmgin"
	"github.com/gin-gonic/gin"
)

// 主要测试：elastic apm
// 运行方式：ELASTIC_APM_SERVER_URL=http://apm-server地址:8200 go run main.go 或可执行程序
func main() {
	r := gin.Default()
	r.Use(apmgin.Middleware(r))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and server on 0.0.0.0:8080
}
