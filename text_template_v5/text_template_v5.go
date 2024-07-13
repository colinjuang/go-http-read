package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmp, err := template.New("test").Parse(`
{{if .yIsZero }}
 除数不能为0
{{else}}
 {{.result}}
{{end}}
`)
		if err != nil {
			fmt.Fprintf(w, "Parse: %v", err)
			return
		}
		x, _ := strconv.ParseInt(r.URL.Query().Get("x"), 10, 64)
		y, _ := strconv.ParseInt(r.URL.Query().Get("y"), 10, 64)

		yIsZero := y == 0
		result := 0.0
		if !yIsZero {
			result = float64(x) / float64(y)
		}

		err = tmp.Execute(w, map[string]interface{}{
			"yIsZero": yIsZero,
			"result":  result,
		})
		if err != nil {
			fmt.Fprintf(w, "Execute: %v", err)
			return
		}
	})
	log.Println("Starting HTTP Server ...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
