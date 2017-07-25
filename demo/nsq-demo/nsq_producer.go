package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/nsqio/go-nsq"
	"log"
	"net/http"
)

type ProBody struct {
	Body    string `json:"body"`    // 主体内容
	Uuid    string `json:"uuid"`    // 时间戳
	Acttime string `json:"acttime"` // 活动时间
	Type    string `json:"type"`    // 类型
	Note    string `json:"note"`    // 备注
}

type ActBody struct {
	Openid  string `json:"openid"`  // 主体内容
	Acttime string `json:"acttime"` // 活动时间
	Type    string `json:"type"`    // 类型
	Note    string `json:"note"`    // 备注
	Uuid    string `json:"uuid"`    // UUID
}

var producer *nsq.Producer

func init() {
	cfg := nsq.NewConfig()
	var err error
	producer, err = nsq.NewProducer("172.16.50.143:4150", cfg)
	//producer, err = nsq.NewProducer("127.0.0.1:4150", cfg)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("nsq 链接正常")
	}
}

func main() {
	e := echo.New()
	e.POST("/producer", func(c echo.Context) error {
		body := new(ProBody)
		var err error
		if err = c.Bind(body); err != nil {
			return c.JSON(http.StatusMethodNotAllowed, "JSON body 内容体不对"+err.Error())
		}
		if body.Body == "" {
			return c.JSON(http.StatusMethodNotAllowed, "body不能为空")
		}

		// 微信活跃数据
		var actBody ActBody
		actBody.Openid = body.Body
		actBody.Acttime = body.Acttime
		actBody.Uuid = body.Uuid
		actBody.Type = body.Type
		actBody.Note = body.Note

		jsons, errs := json.Marshal(actBody) //转换成JSON返回的是byte[]
		if errs != nil {
			fmt.Println(errs.Error())
		}
		if err := producer.Publish("shtelecom-active", jsons); err != nil {
			return c.JSON(http.StatusBadGateway, "publish error: "+err.Error())
		} else {
			e.Logger.Info("SendOk: " + body.Uuid)
			return c.JSON(http.StatusOK, "ok")
		}
	})
	e.Logger.Fatal(e.Start(":8899"))
}
