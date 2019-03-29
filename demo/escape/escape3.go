package main

type S3 struct{}

func ref3(z S3) *S3 {
	return &z
}

func main() {
	var x S3
	_ = *ref3(x)
}
