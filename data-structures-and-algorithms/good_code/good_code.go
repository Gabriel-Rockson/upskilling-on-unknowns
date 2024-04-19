package main

import (
	"fmt"
	"time"
)

func findNemo(array []string) {
	start := time.Now()
	for _, v := range array {
		if v == "nemo" {
			fmt.Println("Found NEMO!!!")
		}
	}
	end := time.Now()
	fmt.Println("The program took ", end.Sub(start))
}

func main() {
	// array := []string{"nemo"}
	bigArray := [100000000000]string{}

	findNemo(bigArray[:])
}
