package main

import "fmt"

type Go struct{}

func (g Go) sayHello() {
	fmt.Println("Hi, I am GO!")
}

type PHP struct{}

func (p PHP) sayHello() {
	fmt.Println("Hi, I am PHP!")
}
func main() {
	golang := Go{}
	php := PHP{}
	golang.sayHello()
	php.sayHello()
}
