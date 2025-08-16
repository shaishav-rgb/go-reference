package main

import (
	"fmt"
)

type house struct {
	name     string
	height   int
	location string
}

func (o house) computeLocation() string {
	return "House"
}

type office struct {
	name     string
	height   int
	location string
}

func (o office) computeLocation() string {
	return "Office"
}

type locationer interface {
	computeLocation() string
}

func compareInterface(o, l locationer) bool {
	return o == l
}

func main() {
	houseValue := house{}
	officeValue := office{}

	fmt.Println(houseValue == house(officeValue))                 //type has to match, gives true as struct is comparable type
	fmt.Println(compareInterface(houseValue, officeValue))        // gives false because types are different
	fmt.Println(compareInterface(houseValue, house(officeValue))) // gives true because interface type and value type both match and struct is a comparable type

	fmt.Println(compareInterface(houseValue, officeValue)) // concrete types are different but allowed in compile, Generics can enforce the same concrete types for these two values that implements interface "locationer" at compile time
}
