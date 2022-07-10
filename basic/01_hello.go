package main

import (
	"fmt"
)

// using println
func hello_v1() {
	println("v1: Hello World from ffr!")
}

// using fmt.Println
func hello_v2() {
	fmt.Println("v2: Hello World from ffr!")
}

func main() {
	hello_v1()
	hello_v2()
}
