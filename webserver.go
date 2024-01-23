package main

import (
	"log"
	"net/http"
)

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
	log.Println("listening..." + s.listenAddr)
	return http.ListenAndServe(s.listenAddr, nil)
}
