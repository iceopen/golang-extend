package main

import (
	"fmt"

	"github.com/cespare/xxhash"
)

func main() {
	data := []byte{1, 2, 3}
	hashVal := xxhash.Sum64(data)
	fmt.Println(xxhash.Sum64(data))
	fmt.Println(xxhash.Sum64String(string(data)))
	slotId := uint8(hashVal >> 8)
	hash16 := uint16(hashVal >> 16)
	fmt.Println(slotId)
	fmt.Println(hash16)
}
