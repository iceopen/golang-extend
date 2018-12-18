package main

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/google/gops/agent"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world")
}

func main() {
	if err := agent.Listen(agent.Options{}); err != nil {
		log.Fatal(err)
	}
	beego.SetLevel(beego.LevelDebug)
	beego.Router("/", &MainController{})
	beego.Run()
}
