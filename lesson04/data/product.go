package data

import (
	"encoding/json"
	"errors"
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

type Products []*Product

func GetProducts() Products {
	return ProductsList
}

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func AddProduct(p *Product) error {
	id := getNextID()
	p.ID = id
	ProductsList = append(ProductsList, p)
	return nil
}

func UpdateProduct(p *Product) error {
	for i, prod := range ProductsList {
		if prod.ID == p.ID {
			ProductsList[i] = p
			return nil
		}
	}

	return errors.New("product not found")
}

func getNextID() uint {
	id := ProductsList[len(ProductsList)-1].ID + 1
	return id
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
