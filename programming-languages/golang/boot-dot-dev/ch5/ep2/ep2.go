package main

import "fmt"

type employee interface {
  getName() string
  getSalary() int
}

type contructor struct {
  name string
  hourlyPay int
  hoursPerYear int
}

func (c contructor) getName() string {
  return c.name
}

func (c contructor) getSalary() int {
  return c.hourlyPay * c.hoursPerYear
}

func main()  {
  firstContructor := contructor{
    name: "Gabriel Rockson",
    hourlyPay: 79,
    hoursPerYear: 2000,
  }

  fmt.Println(firstContructor.getName())
  fmt.Println(firstContructor.getSalary())
}
