package main

import "fmt"

func fact(n int) int {
  if n == 0 {
    return 1
  }

  return n * fact(n - 1)
}

func main() {
  var number int

  fmt.Print("Enter Number: ")
  fmt.Scanln(&number)

  fmt.Println(number, "factorial:", fact(number))
}
