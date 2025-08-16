package main

import (
	"cmp"
	"fmt"
)

type Tree[T any] struct {
	f OrderableFunc[T]
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Order(other Person) int {
	out := cmp.Compare(p.Name, other.Name)
	if out == 0 {
		out = cmp.Compare(p.Age, other.Age)
	}
	return out
}

func OrderPeople(p1, p2 Person) int {
	out := cmp.Compare(p1.Name, p2.Name)
	if out == 0 {
		out = cmp.Compare(p1.Age, p2.Age)
	}
	return out
}

type OrderableFunc[T any] func(t1, t2 T) int

func NewTree[T any](f OrderableFunc[T]) *Tree[T] {
	fmt.Println("Executing")
	return &Tree[T]{
		f: f,
	}
}

func main() {
	var _ = NewTree(OrderPeople)
	_ = NewTree(OrderPeople) // no error, the function executed but no variable is assigned
	// _:= NewTree(OrderPeople)  //cannnot do implicit declare and assign using _
	var _ = NewTree(Person.Order) // using method as a parameter to function type, Person struct becomes the first parameter

}
