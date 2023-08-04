package products

import (
	"math"
)

type Product struct {
	Name  string
	Price float64
}

// FindMostExpensive finds the most expensive product from the list
func FindMostExpensive(products []Product) Product {
	var mostExpensive Product
	maxPrice := math.Inf(-1)

	for _, product := range products {
		if product.Price > maxPrice {
			maxPrice = product.Price
			mostExpensive = product
		}
	}

	return mostExpensive
}
