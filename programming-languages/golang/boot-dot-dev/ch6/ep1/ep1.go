package main

import "fmt"

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (float64, error) {
  costOfMsgToCustomer, err := sendSMS(msgToCustomer)
  if err != nil {
    return 0.0, err
  }

  costOfMsgToSpouse, err := sendSMS(msgToSpouse) 
  if err != nil {
    return 0.0, err
  }

  return costOfMsgToCustomer + costOfMsgToSpouse, nil
}


func sendSMS(message string) (float64, error)  {
  const maxTextLen = 25
  const costPerChar = .0002

  if len(message) > maxTextLen {
    return 0.0, fmt.Errorf("Can't send texts over %v characters", maxTextLen)
  }

  return costPerChar * float64(len(message)), nil
}

func main()  {
  // nothing to do here
}
