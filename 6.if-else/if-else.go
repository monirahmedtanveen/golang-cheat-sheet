package main

import "fmt"

func main() {
  if 3 % 2 == 1 {
    fmt.Println("3 is odd")
  } else {
    fmt.Println("3 is even")
  }

  if 8 % 4 == 0 {
    fmt.Println("8 is divisible by 4")
  } else {
    fmt.Println("8 is not divisible by 4")
  }

  if n := 9; n < 0 {
    fmt.Println(n, "is nagative")
  } else if n < 10 {
    fmt.Println(n, "is a single digit number")
  } else {
    fmt.Println(n, "has multiple digits")
  }

  if true {
    fmt.Println("Hello Golang!!")
  }
}
