package main

import "fmt"

func splitSum(sum int) (x, y int) {
  x = sum * 4 / 9
  y = sum - x
  return
}

func main() {
  x, y := splitSum(20)

  fmt.Println(x, y)
}
