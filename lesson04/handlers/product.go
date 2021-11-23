package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

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

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		p.updateProduct(rw, r)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Product) getProducts(rw http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()
	err := products.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Error on marshal json", http.StatusInternalServerError)
		return
	}
}

func (p *Product) addProduct(rw http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Error on decode json", http.StatusBadRequest)
	}

	data.AddProduct(prod)
}

func (p *Product) updateProduct(rw http.ResponseWriter, r *http.Request) {
	reg := regexp.MustCompile(`/([0-9]+)`)
	g := reg.FindAllStringSubmatch(r.URL.Path, -1)

	if len(g) != 1 {
		http.Error(rw, "Invalid URI", http.StatusBadRequest)
	}

	if len(g[0]) != 2 {
		http.Error(rw, "Invalid URI", http.StatusBadRequest)
	}

	idStr := g[0][1]

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
