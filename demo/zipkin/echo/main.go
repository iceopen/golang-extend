package main

import (
	"io/ioutil"
	"runtime"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/parnurzeal/gorequest"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	e := echo.New()
	e.Debug = false
	e.Use(middleware.BodyLimit("1M"))
	// 处理请求
	e.Any("*", func(c echo.Context) error {
		url := "http://127.0.0.1:9410" + c.Request().URL.Path
		logs.Info("请求地址：", url)
		body, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			return err
		}
		bodyStr := string(body)
		logs.Info("请求内容：", bodyStr)

		request := gorequest.New()
		resp, resBody, _ := request.Timeout(5 * time.Second).Post(url).Send(bodyStr).End()
		logs.Info("zipkin请求返回状态：", resp.StatusCode, "内容：", resBody)
		return c.JSON(resp.StatusCode, "")
	})
	e.Logger.Fatal(e.Start(":9411"))
}
