package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

type Inventory struct {
	SKU       string
	Name      string
	UnitPrice float64
	Quantity  int64
}

func (i *Inventory) Subtotal() float64 {
	return i.UnitPrice * float64(i.Quantity)
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		tmp, err := template.New("test").Parse(`Inventory
Sku: {{.SKU}}
Name: {{.Name}}
UnitPrice: {{.UnitPrice}}
Quantity: {{.Quantity}}
Subtotal: {{.Subtotal}}`)

		if err != nil {
			fmt.Fprintf(writer, "Parse: %v", err)
			return
		}
		inventory := &Inventory{
			SKU:  request.URL.Query().Get("sku"),
			Name: request.URL.Query().Get("name"),
		}

		inventory.UnitPrice, _ = strconv.ParseFloat(request.URL.Query().Get("UnitPrice"), 64)
		inventory.Quantity, _ = strconv.ParseInt(request.URL.Query().Get("Quantity"), 10, 64)

		err = tmp.Execute(writer, inventory)
		if err != nil {
			fmt.Fprintf(writer, "Execute: %v", err)
			return
		}
	})
	log.Println("Starting HTTP server ....")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
