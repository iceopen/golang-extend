package main

import (
	"log"
	_ "net/http/pprof"

	_ "github.com/mkevac/debugcharts"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/gorilla/handlers"
	"fmt"
)

func main() {
	go func() {
		log.Fatal(http.ListenAndServe(":9090", handlers.CompressHandler(http.DefaultServeMux)))
	}()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		urlStr := c.Request.URL.String()
		fmt.Println(urlStr)
		c.JSON(200, gin.H{
			"message": urlStr,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
	log.Printf("you can now open http://localhost:8080/debug/charts/ in your browser")
}
