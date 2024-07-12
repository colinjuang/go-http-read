package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloFunc)
	log.Println("Start http server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
