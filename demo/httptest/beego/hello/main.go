package main

import (
	"github.com/astaxie/beego"
	"github.com/iceopen/igo/pkg/initer"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world")
}

func main() {
	initer.GoPsAgent()
	beego.SetLevel(beego.LevelDebug)
	beego.Router("/", &MainController{})
	beego.Run(":3001")
}
