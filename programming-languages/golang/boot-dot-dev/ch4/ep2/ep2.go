package main

import "fmt"

type messageToSendNew struct {
  message   string
  sender    user
  recipient user
}

type user struct {
  name    string
  number  int
}

func canSendMessage(msgToSend messageToSendNew) bool {

  if msgToSend.sender.name == "" || msgToSend.recipient.name == "" {
    return false
  }

  if msgToSend.sender.number == 0 || msgToSend.recipient.number == 0 {
    return false
  }

  return true
}

func sendMessageNew(message messageToSendNew)  {
  if canSendMessage(message) {
    fmt.Printf("Sending message: '%s' to %v\n", message.message, message.recipient.name)
  }

  fmt.Println("Can't send message, message some key details ...")
}

func main()  {

  messageSender := user{
    name: "Gabriel Rockson",
    number: 679242342,
  }

  messageRecipient := user{
    name: "Handler",
    number: 887722342,
  }

  sendMessageNew(messageToSendNew{
    message: "I couldn't find the target",
    sender: messageSender,
    recipient: messageRecipient,
  })

}
