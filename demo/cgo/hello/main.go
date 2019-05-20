package main

/*
#include <stdio.h>
#include <stdlib.h>

void myprint(char* s) {
	printf("%s\n", s);
}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func Example() {
	cs := C.CString("Hello from stdio\n")
	fmt.Println(cs)
	C.myprint(cs)
	C.free(unsafe.Pointer(cs))
}

func main() {
	Example()
}
