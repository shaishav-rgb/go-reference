package main

import (
	"fmt"
)

type Animal struct {
	Species string
}

func main() {
	var animalPointer interface{} = &Animal{Species: "Cat"}
	var animalValue interface{} = Animal{Species: "Cat"}

	// You must assert to the exact type (*Animal, not Animal)
	a, ok := animalPointer.(*Animal)
	if ok {
		fmt.Println("Species:", a.Species)
	}

	// You must assert to the exact type (Animal, not *Animal)
	b, ok1 := animalValue.(*Animal)
	if ok1 {
		fmt.Println("Species:", b.Species)
	}
}
