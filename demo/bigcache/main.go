package main

import (
	"time"
	"fmt"
	"github.com/allegro/bigcache"
)

func main() {
	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))


	cache.Set("my-unique-key", []byte("value"))

	entry, _ := cache.Get("my-unique-key")
	fmt.Println(string(entry))
	if string(entry) == "" {
		fmt.Println("不存在")
	}
	entry, _ = cache.Get("my-unique-key1")
	if string(entry) == "" {
		fmt.Println("不存在")
	}
	fmt.Println(string(entry))
}
