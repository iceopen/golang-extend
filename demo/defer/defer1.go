package main

import "fmt"

func f1() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func main() {
	fmt.Println(f1())
}
