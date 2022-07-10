package main

import (
	"fmt"
)

func test1() {
	var a, b int
	a = 43
	b = 33
	c := a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)

	x := 32.1
	y := 22.9
	z := x - y
	fmt.Printf("%f - %f = %f\n", x, y, z)
}

func main() {
	test1()
}
