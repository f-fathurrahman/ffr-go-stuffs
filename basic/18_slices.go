package main

import "fmt"

func main() {
	// create an array
	x := make([]float64, 5)
	// populate the array
	for i := 0; i < len(x); i++ {
		x[i] = float64(i+1) * 2.1
	}
	// print them
	fmt.Println(x)
	// using slice
	fmt.Println(x[1:3])
}
