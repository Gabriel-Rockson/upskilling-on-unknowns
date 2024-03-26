package main

import (
	"fmt"
	"time"
)

type birthdayMessage struct {
  birthdayTime time.Time
  recipientName string
}

type report struct {
  reportName string
  numberOfSends int
}

type message interface {
  getMessage() string
}

func sendMessage(m message)  { 
  fmt.Println(m.getMessage())
}

func (birthdayMsg birthdayMessage) getMessage() string  {
  return fmt.Sprintf("Happy birthday %s, your birthday is %s", birthdayMsg.recipientName, birthdayMsg.birthdayTime)
}

func (r report) getMessage() string {
  return fmt.Sprintf("%s is being sent %d number of times", r.reportName, r.numberOfSends)
}

func main()  {
  birthdayMsg := birthdayMessage{
    birthdayTime: time.Now(),
    recipientName: "Gabriel Rockson",
  }

  sampleReport := report{
    reportName: "Exam report",
    numberOfSends: 20,
  }

  sendMessage(birthdayMsg)
  sendMessage(sampleReport)
}
