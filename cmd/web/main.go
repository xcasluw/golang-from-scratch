package main

import "net/http"

const portNumber = ":8080"

func main() {
	http.ListenAndServe(portNumber, nil)
}
