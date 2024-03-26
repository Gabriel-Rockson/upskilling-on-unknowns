package main

import "fmt"

// Type assertions
// We're living in the utopian future where robots are considered
// alive

type living interface {
  getName() string;
  canWalk() bool;
}

type robot struct {
  name string;
}

type animal struct {
  name string;
}

type human struct {
  name string;
}

func (r robot) getName() string {
  return r.name
}

func (a animal) getName() string {
  return a.name
}

func (h human) getName() string {
  return h.name
}

func (r robot) canWalk() bool {
  return true
}

func (a animal) canWalk() bool {
  return true
}

func (h human) canWalk() bool {
  return false
}

func determineLivingBeing(l living) (string, bool) {
  r, ok := l.(robot)
  if ok {
    return r.getName(), r.canWalk()
  }

  return "unidentified being", false
}

func main()  {
  lastHuman :=  human{
    name: "Devin",
  }

  fmt.Println(determineLivingBeing(lastHuman))
}
