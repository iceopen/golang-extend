package main

import (
	"flag"
	"log"

	"github.com/iceopen/golang-extend/project/cache/cache"
	"github.com/iceopen/golang-extend/project/cache/http"
	"github.com/iceopen/golang-extend/project/cache/tcp"
)

func main() {
	cacheType := flag.String("type", "inmemory", "cache type")
	flag.Parse()
	log.Println(cacheType)
	c := cache.New(*cacheType)
	go tcp.New(c).Listen()
	http.New(c).Listen()
}
