package main

import (
	"log"
)

func main() {
	var YES1 string = "yes"
	log.Println(YES1)
	var YES2 string = "yes"
	log.Println(YES2)
	if YES1 == YES2 {
		println("ok")
	}
	// cacheType := flag.String("type", "inmemory", "cache type")
	// node := flag.String("node", "127.0.0.1", "node address")
	// clus := flag.String("cluster", "127.0.0.1", "cluster address")
	// flag.Parse()
	// log.Println(*cacheType)
	// log.Println(*node)
	// log.Println(*clus)
	// c := cache.New(*cacheType)
	// go tcp.New(c).Listen()
	// http.New(c).Listen()
}
