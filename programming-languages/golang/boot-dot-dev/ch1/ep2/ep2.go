package main

import "fmt"

// Fix a bug

func main()  {
  messagesFromDoris := []string{
    "Hi",
    "Good morning!",
    "How are you doing?",
    "I wanted to check out how you're doing given how drunk you were last night",
  }

  numOfMessages := len(messagesFromDoris)
  costPerMessage := .02

  totalCost := costPerMessage * float64(numOfMessages)

  fmt.Printf("Doris spent %.2f on messages.\n", totalCost)
}
