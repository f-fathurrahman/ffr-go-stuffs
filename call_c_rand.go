package main

// The following comment is important !

// #include <stdlib.h>
import "C"

func Random() int {
	return int(C.random())
}

func Seed(i int) {
	C.srandom(C.uint(i))
}

func main() {
	Seed(123)
	println(Random())
}
