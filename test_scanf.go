package main

import (
  "fmt"
  "math"
)

func main() {
  fmt.Println("")
  fmt.Println("Calculating area of a cicle")
  fmt.Println("---------------------------")
  fmt.Println("")
  fmt.Print("Enter a radius of circle: ")

  var radius float64
  fmt.Scanf("%f", &radius)

  area := math.Pi * math.Pow(radius,2)
  fmt.Printf("Area of cicle: %f\n", area)
}
