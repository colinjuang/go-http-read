package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		tmp, err := template.New("test").Parse("the val is {{.}}")
		if err != nil {
			fmt.Fprintf(writer, "Parse: %v", err)
			return
		}
		val := request.URL.Query().Get("val")

		err = tmp.Execute(writer, val)
		if err != nil {
			fmt.Fprintf(writer, "Execute: %v", err)
			return
		}
	})
	log.Println("Starting HTTP server ....")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
