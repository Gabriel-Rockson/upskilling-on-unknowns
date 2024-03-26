package main

import "fmt"

func getSMSErrorString(cost float64, recipient string) string {
  return fmt.Sprint("SMS that costs $%.2f to be sent to '%v' can not be sent", cost, recipient)
}

func test(cost float64, recipient string)  {
  s := getSMSErrorString(cost, recipient)
  fmt.Println(s)
  fmt.Println("==================================")
}
