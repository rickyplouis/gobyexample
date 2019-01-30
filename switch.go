package main

import (
    "fmt";
    "time";
)

func main() {
  switch time.Now().Weekday() {
  case time.Saturday, time.Sunday:
    fmt.Println("It's the weekend!")
  default:
    fmt.Println("It's not the weekend :(")
  }

  whoAmI:= func(i interface{}) {
    switch i.(type) {
    case bool:
      fmt.Println("I'm a bool")
    case int:
      fmt.Println("I'm just another number")
    default:
      fmt.Println("I may never know")
    }
  }

  whoAmI(true)
  whoAmI(2)
  whoAmI("-_-")
}
