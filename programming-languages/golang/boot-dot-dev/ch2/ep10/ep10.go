package main

// Formatting strings
// formatting verbs
// %v - interpolate the default representation
// %s - interpolate a string
// %d - interpolate an integer in decimal form
// %f - interpolate a decimal

import "fmt"

func main() {
  const name = "Saul Goodman"
  const openRate = 30.5

  msg := fmt.Sprintf("Hi %v, your open rate is %.1f percent", name, openRate)

  fmt.Println(msg)
}
