package main

import "fmt"

func main() {
  x := make([]float64, 5)
  for i := 0; i < len(x); i++ {
    x[i] = float64(i+1)*2.1
  }
  fmt.Println(x)
  fmt.Println(x[1:3])
}

