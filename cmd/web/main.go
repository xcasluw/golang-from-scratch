package main

import (
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	server := &http.Server{
		Addr:    portNumber,
		Handler: routes(),
	}
	err := server.ListenAndServe()
	log.Fatal(err)
}
