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
	this.TplName = "index.tpl"
}

func main() {
	if err := agent.Listen(agent.Options{}); err != nil {
		log.Fatal(err)
	}
	beego.Router("/", &MainController{})
	beego.Run()
}
