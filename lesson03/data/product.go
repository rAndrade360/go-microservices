package data

import "time"

type Product struct {
	ID        uint
	Name      string
	Price     float64
	SKU       string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func GetProducts() []*Product {
	return ProductsList
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
