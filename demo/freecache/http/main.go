package main

import (
	"github.com/gin-gonic/gin"
	"github.com/coocood/freecache"
	"runtime/debug"
	"fmt"
)

func main() {
	cacheSize := 100 * 1024 * 1024
	cache := freecache.NewCache(cacheSize)
	debug.SetGCPercent(20)
	key := []byte("abc")
	val := []byte("def")
	expire := 60 // expire in 60 seconds
	cache.Set(key, val, expire)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		got, err := cache.Get(key)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(got))
		}

		fmt.Println("entry count ", cache.EntryCount())
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
