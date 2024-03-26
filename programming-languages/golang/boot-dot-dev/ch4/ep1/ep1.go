package main

import "fmt"

type messageToSend struct {
  phoneNumber int
  message string
}

func sendMessage(message messageToSend)  {
  fmt.Printf("Sending message: '%s' to %v\n", message.message, message.phoneNumber)
}

func main()  {

  sendMessage(messageToSend{
    phoneNumber: 551528489,
    message: "I am enjoying the idea of golang",
  })

}
