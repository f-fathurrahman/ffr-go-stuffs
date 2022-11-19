/* Multiline comment
This is multiline comment
*/


// A package clause starts every source file
// main is a special name declaring an executable rather than a library
package main

import (
	"fmt"
	"reflect"
	"os"
	"math"
)

// A function definition
// main is special. It is the entry point for the executable program.

func main() {
	// Println outputs a line to stdout
	// It comes from the package fmt.
	fmt.Println("Hello world by efefer!")

	// Call another function within this package
	beyondHello()
}

func beyondHello() {
	var x int // variable declaration
	x = 3 // variable assignment

	// "short" declaration, use := to infer the type, declare and assign
	y := 4

	// call a function that returns two values
	ss, pp := learnMultiple(x, y)
	fmt.Println("ss = ", ss, " pp = ", pp)

	learnTypes()
}

func learnMultiple(x int, y int) (ss int, pp int) {
	ss = x + y
	pp = x * y
	// no need to return ss and pp (?)
	return ss, pp
	// Also can be written as: return x + y, x * y
}


func learnTypes() {
	str1 := "My str1"
	str2 := `A "raw" string literal, which can include
line breaks.
Like this`
	// Try using fmt.Printf
	fmt.Printf("str1 = %s\n", str1)
	fmt.Printf("str2 = %s\n", str2)

	// Example unicode
	g := 'Σ' // rune type, an alias for int32, holds a unicode code point
	fmt.Println("g = ", g)

	gg := "Σ" // a string
	fmt.Println("gg = ", gg)
	
	// Using %T in fmt.Printf to get the type
	fmt.Printf("type of str1 = %T\n", str1)
	fmt.Printf("type of str2 = %T\n", str2)
	fmt.Printf("type of g = %T\n", g)
	fmt.Printf("type of gg = %T\n", gg)

	f := 3.14195 // float64
	c := 3 + 4i // complex128
	fmt.Println("f = ", f, " type of f = ", reflect.TypeOf(f))
	fmt.Println("c = ", c, " type of c = ", reflect.TypeOf(c))

	// var syntax with initializers
	var u uint = 8 // unsigned, but implementation dependent size as with int
	fmt.Println("u = ", u, " type of u = ", reflect.TypeOf(u))

	var pi float32 = 22.0/7.0
	fmt.Println("pi = ", pi, " type of pi = ", reflect.TypeOf(pi))

	// Conversion syntax with a short declaration
	n := byte('\n') // byte is an alias for uint8
	fmt.Println("n = ", n, " type of n = ", reflect.TypeOf(n))

	// Arrays have fixed size at compile time
	var a4 [4]int // initialized to all zeros
	fmt.Println("a4 = ", a4)
	fmt.Println("type of a4 = ", reflect.TypeOf(a4))

	aa := [...]int{3, 4, 5, 8, 100, 11}
	fmt.Println("aa = ", aa)
	fmt.Println("type of aa = ", reflect.TypeOf(aa))

	// Arrays have value semantics
	a4_cpy := a4
	// a4_cpy is a copy of a4, two separate instances
	a4_cpy[0] = 9999
	// only a4_cpy is changed, a4 stays the same
	fmt.Println("a4_cpy = ", a4_cpy)
	fmt.Println("a4 = ", a4)

	// Slices have dynamic size.
	s3 := []int{5, 4, 11}
	s4 := make([]int, 4) // allocates slice of 4 ints, initialized to all zeros
	fmt.Println("s3 = ", s3, " type of s3 = ", reflect.TypeOf(s3))
	fmt.Println("s4 = ", s4, " type of s4 = ", reflect.TypeOf(s4))

	var d2 [][]float64 // declaration only, nothing allocated here
	fmt.Println("d2 = ", d2)

	bs := []byte("a slice") // type conversion syntax
	fmt.Println("bs = ", bs, " type of bs = ", reflect.TypeOf(bs))

	// Slices as well as maps and channels have reference semantics
	s3_cpy := s3
	fmt.Println("s3 before = ", s3)
	fmt.Println("s3_cpy before = ", s3_cpy) 
	s3_cpy[1] = 7777
	fmt.Println("s3_cpy after = ", s3_cpy)
	fmt.Println("s3 = after ", s3)

	// Because they are dynamic, slices can be appended to on-demand
	// To append elements to a slide, the built-in append function is used
	// First argument is a slice to which we are appending
	s := []int{3, 4, 5}
	s = append(s, 88, 99, 111) // add 3 elements
	fmt.Println("s = ", s)

	// To append another slice, instead of list of atomic elements we can pass a reference to a slice
	// or a slice literal, with a trailing ellipsis, meaning take a slice and unpack its elements
	// appending them to slice s
	s = append(s, []int{7, 8, 9}...)
	fmt.Println("s = ", s)

	p, q := learnMemory()
	// declares p, q to be type pointer to int
	fmt.Println("*p = ", *p, " type of p = ", reflect.TypeOf(p))
	fmt.Println("*q = ", *q, " type of q = ", reflect.TypeOf(q))

	// Maps are dynamically growable associated array type like hash or dictionary types
	// of some other languages
	m := map[string]int{"three": 3, "four": 4}
	m["one"] = 1
	fmt.Println("m = ", m, " type of m = ", reflect.TypeOf(m))

	// Using _ to ignore value
	file, _ := os.Create("TEMP_output.txt")
	fmt.Fprint(file, "This is how you write to a file, btw")
	file.Close()

	// Next lesson, call a function
	learnFlowControl()
}

// Go is fully garbage collected. It has pointers but no pointers arithmetic.
// You can make a mitake with a nil pointer, but not by incrementing a pointer.
// Unlike in C/C++, taking and returning an address of a local variable is also safe.
func learnMemory() (p, q *int) {
	// named return values p and q have type pointer to int

	p = new(int) // built-in function new allocates memory
	// The allocated int slice is initialized to 0, p is no longer nil

	s := make([]int, 20) // allocate 20 ints as a single block of memory
	s[3] = 7 // assign one of them

	r := -2 // declare another local variable

	return &s[3], &r
	// & takes the address of an object
}

// It is possible unlike many other languages for functions in Go to have named return values.
// Assigning a name to the type being returning in the function declaration line allows us
// to easily return from multiple points in a function as well as to only use the return keyword,
// without anything further
func learnNamedReturns(x, y int) (z int) {
	z = x * y
	return
	// z is implicitly returned here
}

func expensiveComputation() float64 {
	return math.Exp(10)
}

func learnFlowControl() {

	cond := true
	
	// if statement
	if cond {
		fmt.Println("This is inside if")
	}

	cond = false
	if false {
		fmt.Println("This is false")
	} else {
		fmt.Println("This is not false")
	}

	// Use switch in preference to chained if statements
	x := 42.0
	switch x {
	case 0:
	case 1, 2: // can have multiple matches on one case
	case 42:
		// Cases don't fall through
		fmt.Println("This is case 42")
	case 43:
		// not reached
	default:
		fmt.Println("This is case default")
		// default case is optional
	}

	// type switch allows switching on the type of something instead of value
	var data interface{}
	//data = ""
	//data = int64(1234)
	data = 1234
	switch c := data.(type) {
	case string:
		fmt.Println(c, " is a string")
	case int64:
		fmt.Printf("%d is an int64\n", c)
	default:
		// all other cases
		fmt.Println("Default case c = ", c, " type of c = ", reflect.TypeOf(c))
	}

	// Like if, for does not use parens either.
	// Variables declared in for and if are local to their scope
	for x := 0; x < 3; x++ { // ++ is a statement
		fmt.Println("iteration x = ", x)
	}
	fmt.Println("Outside for loop x = ", x, " type of x = ", reflect.TypeOf(x))

	// for is the only loop statement in Go, but it has alternate forms
	for { // infinite loop
		break // break
		continue // unreached
	}

	// You can use range to iterate over an array, a slice, a string, a map, or a channel
	// range returns one (channel) or two values (array, slice, string, and map)
	for key, value := range map[string]int{"one": 1, "two": 2, "three": 3} {
		fmt.Printf("keys = %s value = %d\n", key, value)
	}

	// If you only need the value, use the underscore as the key
	for _, name := range []string{"Bob", "Bill", "Joe"} {
		fmt.Printf("Hello, %s\n", name)
	}

	// As with for, := in an if statement means to declare and assign
	// Calculate y first and then test y > x
	if y := expensiveComputation(); y > x {
		x = y
	}
	fmt.Println("x after if = ", x)

	// Function literals are closures
	xBig := func() bool {
		return x > 10000
		// references x declared above switch statement
	}
	x = 99999
	fmt.Println("xBig = ", xBig())

	x = 1.3e3
	fmt.Println("xBig = ", xBig())

	// Function literals may be defined and called inline, acting as an
	// argument to function, as long as:
	// a) function literal is called immediately
	// b) result type matches expected type of argument
	fmt.Println("Add and double two numbers: ",
		func(a, b int) int {
			return (a + b) * 2
		}(10, 2)) // called with args 10 and 2

	// A goto statement
	goto love

love:
	fmt.Println("This is love")
}
