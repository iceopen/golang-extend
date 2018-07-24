package main

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/message"
	"github.com/astaxie/beego/logs"
)

func replay(ctx *context.Context) {
	//配置微信参数
	config := &wechat.Config{
		AppID:          "",
		AppSecret:      "",
		Token:          "iceinto",
		EncodingAESKey: "",
	}
	wc := wechat.NewWechat(config)

	// 传入request和responseWriter
	server := wc.GetServer(ctx.Request, ctx.ResponseWriter)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
		logs.Info("接收到内容：", msg)
		// 初始化一个返回内容
		var reply *message.Reply
		switch msg.MsgType {
		//文本消息
		case message.MsgTypeText:
			//do something
			text := message.NewText(msg.Content)
			reply = &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
			//图片消息
		case message.MsgTypeImage:
			//do something

			//语音消息
		case message.MsgTypeVoice:
			//do something

			//视频消息
		case message.MsgTypeVideo:
			//do something

			//小视频消息
		case message.MsgTypeShortVideo:
			//do something

			//地理位置消息
		case message.MsgTypeLocation:
			//do something

			//链接消息
		case message.MsgTypeLink:
			//do something

			//事件推送消息
		case message.MsgTypeEvent:
			reply = &message.Reply{message.MsgTypeText, msg.EventKey}
		}
		//
		//回复消息：演示回复用户发送的消息
		//text := message.NewText(msg.Content)
		//return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
		return reply
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
	//发送回复的消息
	server.Send()
}

func main() {
	beego.Any("/", replay)
	beego.Run(":8400")
}
