package main

import "fmt"

type person struct {
  name string
  age int
}

func newPerson(name string, age int) *person {
  p := person{name: name}
  p.name = name
  p.age = age

  return &p
}

func main() {
  fmt.Println(newPerson("John", 48))
  fmt.Println(person{name: "Messi", age: 42})
  fmt.Println(&person{name: "Maisie Young Grandma", age: 23})
}
