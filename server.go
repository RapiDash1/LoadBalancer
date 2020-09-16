package main

import (
	"net/http"
	"net/http/httputil"
)

type server struct {
	URL   string
	proxy *httputil.ReverseProxy
	alive bool
}

func (s *server) CheckHealth() bool {
	res, err := http.Head(s.URL)
	status := s.alive
	if err != nil {
		status = false
		s.alive = false
	} else {
		status = (res.StatusCode == http.StatusOK)
		s.alive = status
	}
	return status
}
