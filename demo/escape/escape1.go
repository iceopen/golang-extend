package main

type S1 struct{}

func identity1(x S1) S1 {
	return x
}

func main() {
	var x S1
	_ = identity1(x)
}
