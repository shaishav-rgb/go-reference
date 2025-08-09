package main

import (
	"fmt"
)

func main() {
	var f *int                           // f is nil
	fmt.Printf("%p is nil pointer\n", f) // f is nil
	var pointerOfPointer = &f
	fmt.Printf("%p address of pointer of pointer\n", pointerOfPointer)
	fmt.Printf("%p is value of pointer of pointer\n", *pointerOfPointer)
	correctUpdateForNilPointer(&f)
	// failedUpdateForNilPointer(f)
	// fmt.Println(*f) // cannot deference a nil pointer, cannot compile compiler panics
	// fmt.Println(f) // prints nil

	// var f *int
	// number := 10
	// f = &number
	// failedUpdateForNonNilPointer(f)
	// correctUpdateForNonNilPointer(f)
	fmt.Println("Outside function")
	fmt.Println(*f)
}

// func failedUpdateForNilPointer(g *int) {
// 	x := 10
// 	g = &x
// }

// func failedUpdateForNonNilPointer(g *int) {
// 	fmt.Println(*g)
// 	x := 100
// 	g = &x
// 	fmt.Println(*g)
// }

func correctUpdateForNilPointer(g **int) {
	fmt.Println("Inside function")
	x := 10
	fmt.Printf("%p is pointer of assigned value\n", &x)
	*g = &x
	fmt.Printf("%p is pointer of pointer passed to function\n", g)
	fmt.Printf("%p is single deference of pointer\n", *g)
	fmt.Printf("%d is value of pointer of pointer inside function\n", **g)
}

func correctUpdateForNonNilPointer(g *int) {
	fmt.Println(*g)
	x := 100
	*g = x
	// g = x //cannot use x (variable of type int) as *int value in assignment
	fmt.Println(*g)
}
