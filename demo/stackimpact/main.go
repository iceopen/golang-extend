package main

import (
	"fmt"
	"net/http"

	"github.com/stackimpact/stackimpact-go"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func main() {
	agent := stackimpact.Start(stackimpact.Options{
		AgentKey: "fa34557e1650b4847f0242d6ece649a99e46fb3e",
		AppName:  "iceinto-demo",
		Debug:    true,
	})

	http.HandleFunc(agent.ProfileHandlerFunc("/", handler))
	http.ListenAndServe(":8080", nil)
}
