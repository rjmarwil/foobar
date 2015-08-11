package main

import (
  "fmt"
)

func main() {
  fmt.Println("Hello World")

  beyondHello()
}

func beyondHello() {
  var x int
  x = 3
  y := 4
  sum, prod := learnMultiple(x, y)
  fmt.Println("sum:", sum, "prod:", prod)

}

func learnMultiple(x, y int) (int, int) {
  return x + y, x * y
}
