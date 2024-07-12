package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &helloHandle{})
	mux.HandleFunc("/timeout", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		w.Write([]byte("time out"))
	})

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		WriteTimeout: 2 * time.Second,
	}
	log.Println("Listening on :8080 ...")
	log.Fatal(server.ListenAndServe())
}

type helloHandle struct{}

func (h *helloHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
