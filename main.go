package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	log.Println("Start http server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
