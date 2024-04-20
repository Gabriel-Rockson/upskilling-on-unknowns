package main

import "fmt"

func logFirstAndSecondBoxes(boxes []int) {
	fmt.Println(boxes[0])
	fmt.Println(boxes[1])
}

func main() {
	boxes := []int{1, 2, 3, 4, 5}

	logFirstAndSecondBoxes(boxes)
}
