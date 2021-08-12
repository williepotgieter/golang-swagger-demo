package main

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
)

func GenerateProducts(n int) []Product {
	var products []Product

	for i := 0; i < n; i++ {
		products = append(products, Product{
			SKU:         uuid.New(),
			Name:        gofakeit.Vegetable(),
			Description: gofakeit.Paragraph(1, 2, 20, ""),
			Price:       gofakeit.Price(1.00, 50.0),
		})
	}

	return products
}
