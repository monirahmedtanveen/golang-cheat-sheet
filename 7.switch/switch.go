package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("Write", i, "as ")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its weekend")
	default:
		fmt.Println("Its a weekday")
	}

  t := time.Now()
	switch {
  case t.Hour() < 12:
    fmt.Println("Its before noon")
  default:
    fmt.Println("Its after noon")
  }

  whatAmI := func(i interface{}) {
    switch t := i.(type) {
    case bool:
      fmt.Println("I'm a bool")
    case int:
      fmt.Println("I'm an int")
    default:
      fmt.Printf("Dont know type %T\n", t)
    }
  }
  whatAmI(false)
  whatAmI(1)
  whatAmI("Golang")
}
