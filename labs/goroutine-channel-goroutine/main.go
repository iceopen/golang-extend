package main

import (
	"fmt"
)

func gen() chan int {
	out := make(chan int)
	fmt.Println("gen start")
	go func() {
		for i := 0; i < 100; i++ {
			out <- i
		}
	}()
	fmt.Println("gen end")
	return out
}

func seq(input chan int) {
	for num := range input {
		fmt.Println(num)
	}
}

func main() {
	in := gen()
	go seq(in)
}
