package main

import "fmt"

// Accumulator 数字计算
func Accumulator() func(int) int {
	var x int
	return func(delta int) int {
		fmt.Printf("(%+v, %+v) - ", &x, x)
		x += delta
		return x
	}
}

func main() {
	var a = Accumulator()
	fmt.Printf("%d\n", a(1))
	fmt.Printf("%d\n", a(10))
	fmt.Printf("%d\n", a(100))

	var b = Accumulator()
	fmt.Printf("%d\n", b(1))
	fmt.Printf("%d\n", b(10))
	fmt.Printf("%d\n", b(100))
}
