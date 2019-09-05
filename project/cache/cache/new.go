package cache

import "log"

// New 创建缓存
func New(typ string) Cache {
	var c Cache
	if typ == "inmemory" {
		c = newInMemoryCache()
	} else if typ == "badger" {
		c = newInBadgerCache()
	}
	if c == nil {
		panic("unknown cache type " + typ)
	}
	log.Println(typ, "ready to serve")
	return c
}
