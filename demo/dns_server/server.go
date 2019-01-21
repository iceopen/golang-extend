package main

import (
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/miekg/dns"
)

type handler struct{}

// ServeDNS 调用
func (this *handler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	msg := dns.Msg{}
	msg.SetReply(r)
	switch r.Question[0].Qtype {
	case dns.TypeA:
		msg.Authoritative = true
		domain := msg.Question[0].Name
		fmt.Println(domain)
		hosts, _ := net.LookupHost(domain)
		fmt.Println(hosts)
		if len(hosts) > 0 {
			for _, address := range hosts {
				msg.Answer = append(msg.Answer, &dns.A{
					Hdr: dns.RR_Header{Name: domain, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 600},
					A:   net.ParseIP(address),
				})
			}
		}
	}
	w.WriteMsg(&msg)
}

func main() {
	srv := &dns.Server{Addr: ":" + strconv.Itoa(53), Net: "udp"}
	srv.Handler = &handler{}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to set udp listener %s\n", err.Error())
	}
}
