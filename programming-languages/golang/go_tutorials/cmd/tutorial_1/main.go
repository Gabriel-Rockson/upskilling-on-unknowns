package main

import "fmt"

func main() {
	intArray := [...]int{1, 2, 3}
	fmt.Println(intArray)

	intSlice1 := []int{1, 2, 3}
	fmt.Printf(
		"Value = %v, Length = %d, Capacity = %d\n",
		intSlice1,
		len(intSlice1),
		cap(intSlice1),
	)
	intSlice1 = append(intSlice1, 9)
	fmt.Printf(
		"Value = %v, Length = %d, Capacity = %d\n",
		intSlice1,
		len(intSlice1),
		cap(intSlice1),
	)

	intSlice2 := []int{12, 13, 15}

	intSlice1 = append(intSlice1, intSlice2...)
	fmt.Printf(
		"Value = %v, Length = %d, Capacity = %d\n",
		intSlice1,
		len(intSlice1),
		cap(intSlice1),
	)

	intSlice3 := make([]int, 0, 50)
	fmt.Printf(
		"Value = %v, Length = %d, Capacity = %d\n",
		intSlice3,
		len(intSlice3),
		cap(intSlice3),
	)
}
