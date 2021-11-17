package handlers

import (
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
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Product) getProducts(rw http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()
	err := data.ToJSON(rw, products)
	if err != nil {
		http.Error(rw, "Error on marshal json", http.StatusInternalServerError)
		return
	}
}
