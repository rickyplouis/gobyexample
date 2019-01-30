package main

import "fmt"

func main() {
  max := 100
  pyth := 0
  tru := 1

  for i := 0; i <= max; i++ {
    if i % 2 != 0 {
      fmt.Println("How odd", i)
    }
  }

  for tru <= max {
    tru = tru + 1
    pyth = pyth + tru
    fmt.Println("pyth", pyth)
  }

}
