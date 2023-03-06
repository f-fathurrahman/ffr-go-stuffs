package main

import "fmt"

func main() {
	
	type Product struct {
		name, category string
		price float64
	}

	// Example of embedded field, it is defined without name
	type StockLevel struct {
		Product
		count int
	}

	stockItem := StockLevel {
		Product: Product { "Kayak", "Watersport", 275.0 },
		count: 100,
	}

	fmt.Println("stockItem = ", stockItem)
	fmt.Println("stockItem.Product = ", stockItem.Product)
}

