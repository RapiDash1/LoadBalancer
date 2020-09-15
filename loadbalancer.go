package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	servers = []string{
		"http://localhost:8000/",
		"http://localhost:8001/",
		"http://localhost:8002/",
		"http://localhost:8003/",
		"http://localhost:8004/",
	}

	lastServedServerIndex = 0
)

func incrementLastServerServerIndex() {
	lastServedServerIndex = (lastServedServerIndex + 1) % len(servers)
}

func serveURL() *url.URL {
	url, _ := url.Parse(servers[lastServedServerIndex])
	incrementLastServerServerIndex()
	return url
}

func loadBalance(res http.ResponseWriter, req *http.Request) {
	url := serveURL()
	revProxy := httputil.NewSingleHostReverseProxy(url)
	revProxy.ServeHTTP(res, req)
}

func main() {
	http.HandleFunc("/", loadBalance)
	http.ListenAndServe(":9090", nil)
}
