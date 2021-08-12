package main

import "github.com/google/uuid"

type Product struct {
	SKU         uuid.UUID `json:"sku"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
}
