package main

import (
	"fmt"
)

func test1() {
	// declare variables
	var mystr string
	var n, m int
	var x float64

	mystr = "This is my string"
	n = 22
	m = 31
	x = 31.4

	fmt.Println("mystr = ", mystr)
	fmt.Println("n     = ", n)
	fmt.Println("m     = ", m)
	fmt.Println("x     = ", x)
}

func test2() {
	var city string = "Bandung"
	var x int = 100
	fmt.Println("city = ", city)
	fmt.Println("x    = ", x)
}

func test3() {
	country := "Indonesia"
	Npoints := 331
	fmt.Println("country = ", country)
	fmt.Println("Npoints = ", Npoints)
}

func test4() {
	var (
		name  string
		email string
		age   int
	)
	name = "Fadjar"
	email = "ffr@qe.com"
	age = 29

	fmt.Println("name  = ", name)
	fmt.Println("email = ", email)
	fmt.Println("age   = ", age)
}

func main() {
	test1()
	test2()
	test3()
	test4()
}
