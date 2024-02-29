package main

import "fmt"

func main() {
  fmt.Println(addTwoNumbers(10, 12))
  fmt.Println(divideTwoNumbers(12, 0))
  fmt.Println(multiplyTwoNumbers(2323, 828))
  fmt.Println(subtractTwoNumbers(89, 79234))
}

func addTwoNumbers(x int, y int) int  {
  return x + y
}

func multiplyTwoNumbers(x, y int) int {
  return x * y
}

func divideTwoNumbers(x, y int) (int, error) {
  if y == 0 {
    return 0, fmt.Errorf("second parameter can't be 0")
  }

  return x / y, nil
}

func subtractTwoNumbers(x int, y int) int {
  return x - y
}
