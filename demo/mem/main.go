package main

import (
	"fmt"
	"runtime"
)

var stat runtime.MemStats

func main() {
	// 获取默认内存
	runtime.ReadMemStats(&stat)
	fmt.Println(stat.HeapSys)

	// 没有初始化，没有初始化都是
	var (
		a struct{}
		b [0]int
		c [100]struct{}
		d = make([]struct{}, 1024)
	)
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &b)
	fmt.Printf("%p\n", &c)
	fmt.Printf("%p\n", &(d[0]))
	fmt.Printf("%p\n", &(d[1]))
	fmt.Printf("%p\n", &(d[1000]))

}
