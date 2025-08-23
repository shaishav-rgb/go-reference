package main

import (
	"fmt"
)

func main() {
	name, age := ignoreNamedReturn()

	fmt.Printf("Name is %v and age is %v\n", name, age)
	returnValue := g()
	fmt.Printf("Return value is %v\n", returnValue)
}

// func ignoreNamedReturn() (name string, age int) { //named return values
// func ignoreNamedReturn() (string, age int) { // cannot name only single variable
func ignoreNamedReturn() (_ string, age int) { // cannot name only single variable so ignoring one named return
	// func ignoreNamedReturn() (string, int) { // unamed return values
	// name = "Panda"
	name1 := "Panda"
	// age:= 23  //cannot redeclare the same variable in this block if named return also has variable with name age, says "no new variables on left side of :="
	// var age= 23 //cannot redeclare the same variable in this block if named return also has variable with name age
	// return

	// return "Tiger", 12
	return name1, 12
	// return name1, age
}

// shadowing can only happen in inner scope not the same outer scope as the name return function
func g() (age int) {
	age = 10 // named return var
	{
		age := 99 // NEW var in inner scope — shadows the outer 'age'
		_ = age   // use inner one(for making go compiler happy)
	}
	return // returns 10 (outer 'age'), not 99
}
