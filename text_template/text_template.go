package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmp, err := template.New("test").Parse("Hello World")
		if err != nil {
			fmt.Fprintf(w, "Parse: %v", err)
			return
		}
		err = tmp.Execute(w, nil)
		if err != nil {
			fmt.Fprintf(w, "Execute: %v", err)
			return
		}
	})
	log.Println("Starting HTTP server ...")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
