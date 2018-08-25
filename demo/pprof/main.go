package main

import (
	"fmt"
	"log"
	"runtime"
	"time"

	_ "net/http/pprof"
	"net/http"
	"runtime/pprof"
	"flag"
	"os"
)

func dummyCPUUsage() {
	var a uint64
	var t = time.Now()
	for {
		t = time.Now()
		a += uint64(t.Unix())
	}
}

func dummyAllocations() {
	var d []uint64

	for {
		for i := 0; i < 2*1024*1024; i++ {
			d = append(d, 42)
		}
		time.Sleep(time.Second * 10)
		fmt.Println(len(d))
		d = make([]uint64, 0)
		runtime.GC()
		time.Sleep(time.Second * 10)
	}
}

func main() {
	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	if *cpuprofile != "" {
		// 根据命令行指定文件名创建 profile 文件
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		// 开启 CPU profiling
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	go dummyAllocations()
	go dummyCPUUsage()
	go func() {
		log.Println(http.ListenAndServe("localhost:8080", nil))
	}()
	log.Printf("you can now open http://localhost:8080/debug/pprof/ in your browser")
	select {}
}
