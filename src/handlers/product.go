package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rAndrade360/go-microservices/data"
)

type Product struct {
	l *log.Logger
}

func NewProductHandler(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) GetProducts(rw http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()
	err := products.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Error on marshal json", http.StatusInternalServerError)
		return
	}
}

func (p *Product) AddProduct(rw http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Error on decode json", http.StatusBadRequest)
	}

	data.AddProduct(prod)
}

func (p *Product) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(rw, "Error on parse string to int", http.StatusBadRequest)
	}

	prod := &data.Product{}

	err = prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Error on decode json", http.StatusBadRequest)
	}

	prod.ID = uint(id)

	data.UpdateProduct(prod)
}
