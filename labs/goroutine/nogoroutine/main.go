package main

import (
	"fmt"
	"math/rand"
	"time"
)

// demo URL http://blog.teamtreehouse.com/goroutines-concurrency
func generateKey() int {
	fmt.Println("Generating key")
	// Super-secret algorithm!
	keys := []int{3, 5, 7, 11}
	key := keys[rand.Intn(len(keys))]
	// It's kinda slow!
	time.Sleep(3 * time.Second)
	fmt.Println("Done generating")
	return key
}

func main() {
	rand.Seed(time.Now().Unix())
	// Call generateKey 3 times.
	for i := 0; i < 3; i++ {
		fmt.Println(generateKey())
	}
	fmt.Println("All done!")
}
