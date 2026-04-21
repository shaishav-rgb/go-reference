package main

import "fmt"

type Animal interface {
	Sound() string
}

type Dog struct{ Name string }

func (d Dog) Sound() string { return "woof" }

type Cat struct{ Name string }

func (c Cat) Sound() string { return "meow" }

func describe(a Animal) {
	// a is an interface — its concrete type could be Dog or Cat

	if d, ok := a.(Dog); ok {
		fmt.Println("It's a dog named", d.Name)
	} else if c, ok := a.(Cat); ok {
		fmt.Println("It's a cat named", c.Name)
	}
}

func main() {
	describe(Dog{Name: "Rex"}) // asserts through Animal interface
	describe(Cat{Name: "Whiskers"})

	// d := Dog{Name: "Rex"}
	// _,isDog := d.(Dog)  // ← compile error: Dog is not an interface
	// fmt.Println(isDog)
}
