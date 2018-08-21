package gorequest

import (
	"testing"
	"github.com/parnurzeal/gorequest"
	"github.com/stretchr/testify/assert"
	"time"
)

func TestPostJSON(t *testing.T) {
	// http://127.0.0.1:8080/post
	// http://httpbin.org/post
	request := gorequest.New()
	resp, body, errs := request.Timeout(3 * time.Second).Post("http://127.0.0.1:8080/post").
		Send(`{"table":"dx_weixin_event","channel":"WX","data":"{\"ToUserName\":\"\",\"FromUserName\":\"\",\"CreateTime\":\"1534819335\",\"MsgType\":\"event\",\"Event\":\"CLICK\",\"EventKey\":\"手机专区\"}"}`).
		End()
	assert.NotEmpty(t, errs, "请求报错")
	assert.Empty(t, body, "请求结果错误")
	assert.Empty(t, resp, "请求处理错误")
	t.Log(body)
}
