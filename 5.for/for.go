package main

import "fmt"

func main() {
  i := 1
  for i <= 3 {
    fmt.Print(i, " ")
    i++
  }
  fmt.Println()

  for j := 1; j < 9; j++ {
    fmt.Print(j, " ")
  }
  fmt.Println()

  for {
    fmt.Println("Loop Break")
    break
  }

  for n := 1; n < 10; n++ {
    if n % 2 == 0 {
      continue
    }
    fmt.Print(n, " ")
  }
  fmt.Println()
}
