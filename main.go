package main

import (
	"log"
	"net/http"
	"google.golang.org/appengine"
)

func main() {
	log.Println("Booting up content-creator")
	if router := buildRouter(); router != nil {
		http.Handle("/", router)
	}
	appengine.Main()
}
