package main

import (
	"log"
	_ "net/http/pprof"

	_ "github.com/mkevac/debugcharts"
	"github.com/astaxie/beego"
	"net/http"
	"github.com/gorilla/handlers"
)

func main() {
	go func() {
		log.Fatal(http.ListenAndServe(":9090", handlers.CompressHandler(http.DefaultServeMux)))
	}()

	log.Printf("you can now open http://localhost:8080/debug/charts/ in your browser")
	beego.Run()
}
