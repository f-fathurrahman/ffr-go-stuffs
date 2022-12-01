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

	var lifejacket Product
	fmt.Println("lifejacket.name = ", lifejacket.name)
	fmt.Println("lifejacket.category = ", lifejacket.category)
	fmt.Println("lifejacket.price = ", lifejacket.price)
	
}

