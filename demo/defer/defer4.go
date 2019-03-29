package main

import (
	"errors"
	"fmt"
)

func e1() {
	var err error
	defer fmt.Println(1, err)
	err = errors.New("defer1 error")
	return
}
func e2() {
	var err error
	defer func() {
		fmt.Println(2, err)
	}()
	err = errors.New("defer2 error")
	return
}
func e3() {
	var err error
	defer func(err error) {
		fmt.Println(3, err)
	}(err)
	err = errors.New("defer13 error")
	return
}

func main() {
	e1()
	e2()
	e3()
}
