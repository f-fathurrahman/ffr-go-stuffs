package main

import (
  "fmt"
  "runtime"
)

func main() {
  fmt.Printf("This Go version: %s\n", runtime.Version())
}
