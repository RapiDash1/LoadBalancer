package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func createReverserProxy(serverURLString string) *httputil.ReverseProxy {
	serverURLgo, _ := url.Parse(serverURLString)
	return httputil.NewSingleHostReverseProxy(serverURLgo)
}

func createServer(serverURLString string) server {
	return server{serverURLString, createReverserProxy(serverURLString), false}
}

var (
	servers = []server{
		createServer("http://localhost:8000/"),
		createServer("http://localhost:8001/"),
		createServer("http://localhost:8002/"),
		createServer("http://localhost:8003/"),
		createServer("http://localhost:8004/"),
	}

	lastServedServerIndex = 0
)

func incrementLastServerServerIndex() {
	lastServedServerIndex = (lastServedServerIndex + 1) % len(servers)
}

func serveURL() (*httputil.ReverseProxy, error) {
	for i := 0; i < len(servers); i++ {
		incrementLastServerServerIndex()
		if servers[lastServedServerIndex].alive {
			return servers[lastServedServerIndex].proxy, nil
		}
	}
	return nil, fmt.Errorf("No servers available")
}

func loadBalance(res http.ResponseWriter, req *http.Request) {
	revProxy, err := serveURL()
	if err != nil {
		res.Write([]byte(fmt.Sprint(err)))
		return
	}
	revProxy.ServeHTTP(res, req)
}

func main() {
	http.HandleFunc("/", loadBalance)
	go healthCheckAllServers()
	http.ListenAndServe(":9090", nil)
}
