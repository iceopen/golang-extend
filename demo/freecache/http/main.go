package main

import (
	"github.com/gin-gonic/gin"
	"github.com/coocood/freecache"
	"runtime/debug"
	"fmt"
	"time"
)

func main() {
	cacheSize := 100 * 1024 * 1024
	cache := freecache.NewCache(cacheSize)
	debug.SetGCPercent(20)
	key := []byte("abc")
	val := []byte("def")
	expire := 10 // expire in 60 seconds
	cache.Set(key, val, expire)
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		got, err := cache.Get(key)
		if err != nil {
			val = []byte(fmt.Sprint(time.Now().Unix()))
			cache.Set(key, val, expire)
			fmt.Println(string(got))
		} else {
			//fmt.Println(string(got))
		}
		c.JSON(200, gin.H{
			"message": string(got),
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
