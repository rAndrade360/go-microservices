package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	SKU       string    `json:"sku"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `json:"-"`
}

func GetProducts() []*Product {
	return ProductsList
}

func ToJSON(w io.Writer, p []*Product) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

var ProductsList = []*Product{
	{
		ID:        1,
		Name:      "Coffee",
		Price:     27.5,
		SKU:       "1234",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now().Add(time.Hour * 2),
	},
	{
		ID:        2,
		Name:      "Milk",
		Price:     12.8,
		SKU:       "5412",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now().Add(time.Hour * 2),
	},
}
