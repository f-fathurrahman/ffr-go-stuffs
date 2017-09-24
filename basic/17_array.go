package main

import "fmt"

func test1() {
  var x [5]int
  x[1] = 100
  fmt.Println(x)
}

func test2() {
  var x [5]float64
  x[0] = 98
  x[1] = 33
  x[2] = 45
  x[3] = 77
  x[4] = 92

  var total float64 = 0
  for i := 0; i < len(x); i++ {
    total = total + x[i]
  }
  fmt.Println("x = ", x)
  fmt.Println("average = ", total/float64(len(x)))
}

func test3() {
  x := [5]float64{ 98, 70, 81, 77, 65 }
  fmt.Println(x)
  // using another form of loop, looping var is not used
  // single underscore tell the compiler that we don't need this (iterator variable)
  total := 0.0
  for _, value := range x {
    total = total + value
  }
  fmt.Println(total/float64(len(x)))
}

func main() {
  test1()
  test2()
  test3()
}
