package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server")
	log.Fatal(http.ListenAndServe(":8080", &helloHandler{}))
}

type helloHandler struct{}

func (h helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
