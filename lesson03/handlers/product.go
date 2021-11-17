package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rAndrade360/go-microservices/data"
)

type Product struct {
	l *log.Logger
}

func NewProductHandler(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()
	d, err := json.Marshal(products)
	if err != nil {
		http.Error(rw, "Error on marshal json", http.StatusInternalServerError)
		return
	}

	rw.Write(d)
}
