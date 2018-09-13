package httplib

import (
	"testing"
	"time"
	"github.com/astaxie/beego/httplib"
)

type BodyModel struct {
	Table   string `json:"table"`
	Channel string `json:"channel"`
	Data    string `json:"data"`
}

func TestHttplib(t *testing.T) {
	body := BodyModel{}
	body.Channel = "WX"
	body.Data = "{\"ToUserName\":\"gh_559d630d72ec\",\"FromUserName\":\"oKXUCj6Np6hwdre2X1Ldp3fLw7IA\",\"CreateTime\":\"1534819335\",\"MsgType\":\"event\",\"Event\":\"CLICK\",\"EventKey\":\"手机专区\"}"
	body.Table = "dx_weixin_event"
	//ConsumerStart("127.0.0.1:4150", "sjy")
	// "http://httpbin.org/post"
	// http://httpbin.org/post
	var req, err = httplib.Post("http://127.0.0.1:8080/testFlumeHttp").SetTimeout(3*time.Second, 3*time.Second).JSONBody(body)
	//req.Header("Content-Type","application/json;charset=utf-8")
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	t.Log(req.String())
}
