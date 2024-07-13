package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmp, err := template.New("test").Parse(`
{{$name := "alice"}}
{{$age := "18"}}
{{$round2 := true}}
Name: {{$name}}
Age: {{$age}}
Round2: {{$round2}}

{{$name = "bob"}}
Name: {{$name}}

`)
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
	log.Println("Starting HTTP Server ...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
