package main

import "fmt"

func intSec() func() int {
  i := 5
  return func() int {
    i++
    return i + 2
  }
}

func main() {
  nextInt := intSec()

  fmt.Println(nextInt())
  fmt.Println(nextInt())
  fmt.Println(nextInt())

  newInts := intSec()

  fmt.Println(newInts())
}
