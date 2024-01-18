package main

import "log"

func main() {
	app := NewWebserver(":9000")
	log.Fatal(app.Run())
}
