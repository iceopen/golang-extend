package main

import (
	"fmt"

	"github.com/tidwall/evio"
)

// func main() {
// 	listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("10.253.3.22"), Port: 5000})
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Printf("Local: <%s> \n", listener.LocalAddr().String())
// 	data := make([]byte, 8096)
// 	for {
// 		n, remoteAddr, err := listener.ReadFromUDP(data)
// 		if err != nil {
// 			fmt.Printf("error during read: %s", err)
// 		}
// 		fmt.Println("The buf len:", n, remoteAddr)
// 		//fmt.Println(data[:n])
// 		//cs := C.CString(data[:n])
// 		//defer C.free(unsafe.Pointer(cs))
// 		//fmt.Println(cs)
// 	}
// }

// func main() {
// 	var port int
// 	var loops int
// 	var udp bool
// 	var trace bool
// 	var reuseport bool
// 	var stdlib bool

// 	flag.IntVar(&port, "port", 5000, "server port")
// 	flag.BoolVar(&udp, "udp", false, "listen on udp")
// 	flag.BoolVar(&reuseport, "reuseport", false, "reuseport (SO_REUSEPORT)")
// 	flag.BoolVar(&trace, "trace", false, "print packets to console")
// 	flag.IntVar(&loops, "loops", 0, "num loops")
// 	flag.BoolVar(&stdlib, "stdlib", false, "use stdlib")
// 	flag.Parse()

// 	var events evio.Events
// 	events.NumLoops = loops
// 	events.Serving = func(srv evio.Server) (action evio.Action) {
// 		log.Printf("echo server started on port %d (loops: %d)", port, srv.NumLoops)
// 		if reuseport {
// 			log.Printf("reuseport")
// 		}
// 		if stdlib {
// 			log.Printf("stdlib")
// 		}
// 		return
// 	}
// 	events.Data = func(c evio.Conn, in []byte) (out []byte, action evio.Action) {
// 		if trace {
// 			log.Printf("%s", strings.TrimSpace(string(in)))
// 		}
// 		out = in
// 		log.Panicln(in)
// 		return
// 	}
// 	scheme := "tcp"
// 	if udp {
// 		scheme = "udp"
// 	}
// 	if stdlib {
// 		scheme += "-net"
// 	}
// 	log.Fatal(evio.Serve(events, fmt.Sprintf("%s://:%d?reuseport=%t", scheme, port, reuseport)))
// }
func main() {
	var events evio.Events
	events.Data = func(c evio.Conn, in []byte) (out []byte, action evio.Action) {
		out = in
		fmt.Println(out)
		return
	}
	if err := evio.Serve(events, "tcp://localhost:5000"); err != nil {
		panic(err.Error())
	}
}
