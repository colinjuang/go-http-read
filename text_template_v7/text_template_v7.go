package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Inventory struct {
	SKU       string
	Name      string
	UnitPrice float64
	Quantity  int64
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		//		tmp, err := template.New("test").Parse(`Inventory
		//Sku: {{.Inventory.SKU}}
		//Name: {{.Inventory.Name}}
		//UnitPrice: {{.Inventory.UnitPrice}}
		//Quantity: {{.Inventory.Quantity}}`)
		tmp, err := template.New("test").Parse(`Inventory
{{- with .Inventory}}
Sku: {{.SKU}}
Name: {{.Name}}
UnitPrice: {{.UnitPrice}}
Quantity: {{.Quantity}}
{{- end}}`)

		if err != nil {
			fmt.Fprintf(writer, "Parse: %v", err)
			return
		}

		err = tmp.Execute(writer, map[string]interface{}{
			"Inventory": &Inventory{
				SKU:       "1232",
				Name:      "苹果",
				UnitPrice: 0.55,
				Quantity:  88,
			},
		})
		if err != nil {
			fmt.Fprintf(writer, "Execute: %v", err)
			return
		}
	})
	log.Println("Starting HTTP server ....")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
