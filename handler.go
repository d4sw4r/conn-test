package main

import (
	"html/template"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		return
	}
	tmpl.Execute(w, nil)

}

func testHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url := r.FormValue("url")
	check := r.FormValue("check")
	tmpl, _ := template.ParseFiles("template/index.html")
	m := make(map[string]string)
	m["kind"] = "success"
	m["msg"] = url + " successful tested!"
	var err error

	if check == "url" {
		err = HttpTest(url)
	}
	if check == "icmp" {
		err = IcmpTest(url)
	}
	if err != nil {
		m["kind"] = "danger"
		m["msg"] = err.Error()

	}
	tmpl.Execute(w, m)

}
