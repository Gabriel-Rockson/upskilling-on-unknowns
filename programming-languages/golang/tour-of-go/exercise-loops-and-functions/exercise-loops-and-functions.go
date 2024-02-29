package main

import "fmt"

func Sqrt(x float64) float64 {
	z := 1.0
	prevValue := 0.0

	for {
		z -= (z*z - x) / (2 * z)

		if prevValue == z {
			return z
		}

		prevValue = z

		fmt.Printf("prevValue = %v, z = %v\n", prevValue, z)
	}
}

func main() {
	fmt.Println(Sqrt(2))
}
