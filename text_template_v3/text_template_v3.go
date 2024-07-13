package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmp, err := template.New("test").Parse(`{{/*打印参数的值*/}}
Inventory
SKU: {{.SKU}}
Name: {{.Name}}
UnitPrice: {{.UnitPrice}}
Quantity: {{.Quantity}}`)
		if err != nil {
			fmt.Fprintf(w, "Parse: %v", err)
			return
		}

		sku := r.URL.Query().Get("sku")
		name := r.URL.Query().Get("name")
		unitPrice := r.URL.Query().Get("UnitPrice")
		quantity := r.URL.Query().Get("Quantity")

		err = tmp.Execute(w, map[string]interface{}{
			"SKU":       sku,
			"Name":      name,
			"UnitPrice": unitPrice,
			"Quantity":  quantity,
		})
		if err != nil {
			fmt.Fprintf(w, "Execute: %v", err)
			return
		}
	})
	log.Println("Starting HTTP Server ...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
