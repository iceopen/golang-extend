package main

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world")
}

func main() {
	beego.SetLevel(beego.LevelDebug)
	beego.Router("/", &MainController{})
	beego.Run()
}
