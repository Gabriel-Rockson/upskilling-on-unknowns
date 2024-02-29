package main

import "fmt"

func swap(firstName, lastName string) (string, string)  {
  return lastName, firstName
}

func main()  {
  fmt.Println(swap("Gabriel", "Rockson"))
}
