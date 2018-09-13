package main

import (
	"testing"
	"github.com/coocood/freecache"
	"runtime/debug"
		)

func BenchmarkTest(b *testing.B) {
	cacheSize := 100 * 1024 * 1024
	cache := freecache.NewCache(cacheSize)
	debug.SetGCPercent(20)
	key := []byte("abc")
	val := []byte("def")
	expire := 6000 // expire in 60 seconds
	cache.Set(key, val, expire)

	for i := 0; i < b.N; i++ {
		_, err := cache.Get(key)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestTTT(t *testing.T) {



}
