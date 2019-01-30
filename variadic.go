package main

import (
    "fmt"
)

func main() {
  nums := []int{1, 2, 3, 4, 5}
  fmt.Println("nums", nums)
  sum(nums...)
}

func sum(nums ...int) {
  total := 0
  for _, num := range nums {
    total += num
  }
  fmt.Println(total)
}
