package main

import "net/http"

type Webserver struct {
	listenAddr string
}

func NewWebserver(listenaddr string) *Webserver {
	return &Webserver{
		listenAddr: listenaddr,
	}
}
func (s *Webserver) Run() error {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/test", testHandler)
	return http.ListenAndServe(s.listenAddr, nil)
}
