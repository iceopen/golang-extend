package main

import (
	"fmt"

	lru "github.com/hashicorp/golang-lru"
)

func main() {
	l, _ := lru.New(128)
	for i := 0; i < 256; i++ {
		l.Add(i, "ok")
	}
	if l.Len() != 128 {
		panic(fmt.Sprintf("bad len: %v", l.Len()))
	}
	fmt.Println(l.Len())

	fmt.Println(l.Contains(1))
	fmt.Println(l.Len())
}
