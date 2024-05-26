package main

import (
	"errors"
	"fmt"
)

func main() {
	var myValue string = "Hello, World! I am coding in Golang"
	printMe(myValue)

	var numerator int = 11
	var denominator int = 5

	result, remainder, err := intDivision(numerator, denominator)
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Printf("Remainder = %d, Result = %d\n", remainder, result)
	}

	switch remainder {
	case 0:
		fmt.Println("The division was exact")
	case 1, 2:
		fmt.Println("There was just 1 or 2 overflow")
	default:
		fmt.Println("There was a lot of overflow")
	}
}

func printMe(printVal string) {
	fmt.Println(printVal)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("Denominator cannot be 0")
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator

	return result, remainder, nil
}
