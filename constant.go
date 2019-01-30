package main

import "fmt"
import "math"

const s string = "truuuu"

func main() {
  fmt.Println(s)
  const test = 1 + 2
  const another = 3.0 / 7.2
  fmt.Println(test)
  fmt.Println(another)
  fmt.Println(test)
  fmt.Println(math.Sin(test))
}
