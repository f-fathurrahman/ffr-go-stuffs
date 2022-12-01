package main

import "fmt"

func main() {
	
	type Product struct {
		name, category string
		price float64
	}

	kayak := Product {
		name: "Kayak",
		category: "Watersports",
	}
	// price is not assigned
	// price will be assigned to "zero" value
	
	fmt.Println("kayak = ", kayak)
	fmt.Println(kayak.name, kayak.category, kayak.price)

	kayak.price = 300
	fmt.Println("Changed price: ", kayak.price)
		
}

