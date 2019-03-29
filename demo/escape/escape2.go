package main

type S2 struct{}

func identity2(x *S2) *S2 {
	return x
}

func main() {
	var x S2
	y := &x
	_ = identity2(y)
}
