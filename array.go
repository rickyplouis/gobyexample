package main

import (
    "fmt"
)

func main() {
    var ayee [5] int
    var bay[100] int
    kay := [5]int{1, 2, 3, 4, 5}

    for i := 0; i < 100; i++ {
      bay[i] = i
    }

    var test[10][10]int
    for j := 0; j < len(test[0]); j++ {
      for k := 0; k < len(test[1]); k++ {
        test[j][k] = j + k
      }
    }
    fmt.Println(ayee)
    fmt.Println(bay)
    fmt.Println(kay)
    fmt.Println(test)
}
