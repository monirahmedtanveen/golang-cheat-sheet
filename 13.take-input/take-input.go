package main

import "fmt"

func main() {
  var first_name, last_name string

  fmt.Println("First Name:")
  fmt.Scanln(&first_name)

  fmt.Println("Last Name:")
  fmt.Scanln(&last_name)

  fmt.Println("Your Name:", first_name, last_name)
}
