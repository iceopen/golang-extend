package main

import (
	_ "iceinto/apiproject/routers"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego"
	"fmt"
)

func main() {
	fmt.Println("当前运行模式：" + beego.BConfig.RunMode)
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	logs.EnableFuncCallDepth(true)
	logs.Async(1e3)
	logs.SetLogger("console")
	logs.SetLevel(7)
	logs.SetLogger(logs.AdapterFile,`{"filename":"logs/project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":30}`)

	beego.Run()
}
