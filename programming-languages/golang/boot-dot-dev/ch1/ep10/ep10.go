package main

import "fmt"

func main()  {
  username := "wagtail"
  password := "123456789"

  fmt.Println("Authorization: Basic", username + ":" + password)
}
