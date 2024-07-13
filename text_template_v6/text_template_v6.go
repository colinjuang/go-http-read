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
{{$name1 := "alice"}}
{{$age1 := 18}}
{{$name2 := "bob"}}
{{$age2 := 24}}

{{if eq $name1 $name2}}
 名字相同
{{else}}
 名字不相同
{{end}}

{{if eq $age1 $age2}}
 年龄相同
{{else}}
 年龄不相同
{{end}}

{{if gt $age1 $age2}}
	alice 年龄大
{{else}}
	bob 年龄大
{{end}}

{{range $name, $val := .}}
	{{$name}} {{$val}}
{{end}}
`)
		if err != nil {
			fmt.Fprintf(w, "Parse: %v", err)
			return
		}

		err = tmp.Execute(w, map[string]interface{}{
			"Names":   []string{"alice", "bob", "mike", "john"},
			"Numbers": []int{1, 23, 34, 54, 66},
		})
		if err != nil {
			fmt.Fprintf(w, "Execute: %v", err)
			return
		}
	})
	log.Println("Starting HTTP Server ...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
