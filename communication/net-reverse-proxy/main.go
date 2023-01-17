package main

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func main() {
	go proxyServer()
	go targetServer()

	time.Sleep(1 * time.Second)

	for i := 0; i < 10; i++ {
		http.Post("http://127.0.0.1:10000/test", "application/json", nil)
		time.Sleep(100 * time.Millisecond)
	}

	for i := 0; i < 2; i++ {
		http.Post("http://127.0.0.1:10000/other", "application/json", nil)
		time.Sleep(100 * time.Millisecond)
	}

	select {}
}

func proxyServer() {
	target, _ := url.Parse("http://127.0.0.1:20000")

	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ModifyResponse = func(r *http.Response) error {
		fmt.Printf("from proxy %s code=%#v\n", r.Request.URL.Path, r.StatusCode)
		return nil
	}

	http.ListenAndServe("127.0.0.1:10000", proxy)
}

func targetServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("**target handle\n")
		w.Write([]byte("Hello from test"))
	})

	lis, _ := net.Listen("tcp", "127.0.0.1:20000")
	http.Serve(lis, mux)
}
