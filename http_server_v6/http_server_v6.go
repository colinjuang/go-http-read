package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	mex := http.NewServeMux()
	mex.Handle("/", &helloHandler{})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mex,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	go func() {
		<-quit

		if err := server.Close(); err != nil {
			log.Fatal("close server:", err)
		}
	}()

	log.Println("Starting HTTP server...")
	err := server.ListenAndServe()
	if err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			log.Println("server closed under construction")
		} else {
			log.Fatal("server error:", err)
		}
	}
}

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
