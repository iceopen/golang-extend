package main

import "fmt"

func f3() (r int) {
	defer func(r *int) {
		*r = *r + 5
	}(&r)
	return r
}

func main() {
	fmt.Println(f3())
}
