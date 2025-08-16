package main

import (
	"fmt"
)

type Thinger interface {
	Thing()
}
type ThingerInt int

func (t ThingerInt) Thing() {
	fmt.Println("ThingInt:", t)
}

type ThingerSlice []int

func (t ThingerSlice) Thing() {
	fmt.Println("ThingSlice:", t)
}

func Comparer[T comparable](t1, t2 T) {
	if t1 == t2 {
		fmt.Println("equal!")
	}
}



func main() {
	var a int = 10
	var b int = 10
	Comparer(a, b) // prints true
	var a2 ThingerInt = 20
	var b2 ThingerInt = 20
	Comparer(a2, b2) // prints true

	var a3 ThingerSlice = []int{1, 2, 3}
	var b3 ThingerSlice = []int{1, 2, 3}
	// Comparer(a3, b3) // compile fails: "ThingerSlice does not satisfy comparable"

	var a4 Thinger = a2 // type of a variable can be interface(only interface with type elements cannot be a variable type. interfaces with type elements are valid only as type constraints)
	var b4 Thinger = b2
	Comparer(a4, b4) // prints true

	a4 = a3
	b4 = b3
	Comparer(a4, b4) //panics as a4 and b4 are interface but underlying type is slice
}
