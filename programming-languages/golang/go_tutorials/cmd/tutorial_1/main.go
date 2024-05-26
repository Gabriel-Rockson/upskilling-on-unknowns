package main

import (
	"fmt"
	"strings"
)

func main() {
	myString := "sumário"
	fmt.Println(len(myString))
	fmt.Printf("%v, %T\n", myString, myString)

	for i, v := range myString {
		fmt.Printf("%d, %v\n", i, v)
	}

	anotherString := []rune("sumário")
	for i, v := range anotherString {
		fmt.Printf("%v, %v\n", i, v)
	}

	// building a string

	// this is an inefficient way to do it
	myNameChars := []string{"G", "a", "b", "r", "i", "e", "l"}
	nameString := ""
	for i := range myNameChars {
		nameString += myNameChars[i]
	}

	fmt.Println(nameString)

	// this is an efficient way to do it
	anotherNameSlice := []string{"M", "e", "l", "i", "n", "d", "a"}
	var stringBuilder strings.Builder
	for i := range anotherNameSlice {
		stringBuilder.WriteString(anotherNameSlice[i])
	}
	name := stringBuilder.String()
	fmt.Println(name)
}
