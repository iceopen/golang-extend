package main

import (
	"net"
	"os"
)

func main() {
	//Connect udp
	conn, err := net.Dial("tcp", "127.0.0.1:5000")
	if err != nil {
		os.Exit(1)
	}
	defer conn.Close()

	//simple Read
	buffer := make([]byte, 1024)
	conn.Read(buffer)

	//simple write
	conn.Write([]byte("Hello from client"))

}
