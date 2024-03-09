package main

import "fmt"

func personCapacilitiesAccordingToAge(age int) {
	switch age {
	case 10, 11, 12, 13, 14, 15:
		fmt.Println("You are too young to even go out alone")
	case 18, 19, 20, 21:
		fmt.Println("You can have a partner now, Yipee!!!")
	default:
		fmt.Println("What you are capable of will be added soon")
	}

}

func takeUserAge() (userAge int) {
	fmt.Println("How old are you? ")
	var age int

	_, err := fmt.Scanln(&age)
	if err != nil {
		return
	}

	return
}

func main() {
	userAge := takeUserAge()
	personCapacilitiesAccordingToAge(userAge)
}
