package main

import "fmt"

func main() {
	// Maps in golang
	firstMap := make(map[string]uint)
	fmt.Printf("Value = %v, Length = %d\n", firstMap, len(firstMap))
	firstMap["Gabriel"] = 24
	fmt.Printf("Value = %v, Length = %d\n", firstMap, len(firstMap))

	secondMap := map[string]uint{"Gabriel": 24, "Melinda": 22}
	fmt.Printf("Value = %v\n", secondMap)

	age, ok := secondMap["Melindas"]
	if !ok {
		fmt.Printf("Name Melinda is not in the map\n")
	} else {
		fmt.Printf("Age of Melinda is %v\n", age)
	}

	// Looping

	// looping over a map ( the order is not preserved when looping over maps )
	for name, age := range secondMap {
		fmt.Printf("Name = %v, Age = %d\n", name, age)
	}

	// looping over an array or slice
	nameSlice := []string{"Gabriel", "Rockson", "Melinda", "Ampah", "Korsah"}
	for i, v := range nameSlice {
		fmt.Printf("Index = %d, Name = %v\n", i, v)
	}
}
