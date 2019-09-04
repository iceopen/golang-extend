package main

import (
	"github.com/iceopen/golang-extend/project/cache/cache"
	"github.com/iceopen/golang-extend/project/cache/http"
	"github.com/iceopen/golang-extend/project/cache/tcp"
)

func main() {
	c := cache.New("inmemory")
	go tcp.New(c).Listen()
	http.New(c).Listen()
}
