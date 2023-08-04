package main

import (
	"fmt"

	"github.com/kishore-tadapaneni/productlist/products"
)

func main() {
	// Sample list of products
	productList := []products.Product{
		{Name: "Product A", Price: 25.99},
		{Name: "Product B", Price: 19.99},
		{Name: "Product C", Price: 35.50},
		{Name: "Product D", Price: 42.75},
		{Name: "Product E", Price: 29.99},
	}

	// Find the most expensive product
	mostExpensive := products.FindMostExpensive(productList)

	// Display details of the most expensive product
	fmt.Printf("Most Expensive Product:\n")
	fmt.Printf("Name: %s\n", mostExpensive.Name)
	fmt.Printf("Price: र%.2f\n", mostExpensive.Price)
}
