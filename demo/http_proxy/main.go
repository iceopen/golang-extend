package main

import (
	"net/http"
	"net/http/httputil"
	"log"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		director := func(req *http.Request) {
			req = r
			req.URL.Scheme = "http"
			//req.URL.Scheme = r.URL.Scheme
			req.URL.Host = r.Host
		}
		log.Println(r.UserAgent())
		log.Println(r.RequestURI)
		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(w, r)
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}
