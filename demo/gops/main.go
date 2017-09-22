package main

import (
	"log"
	"time"

	"github.com/google/gops/agent"
)

func main() {
	if err := agent.Listen(nil); err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Hour)
}
