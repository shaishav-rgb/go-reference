package main

import "fmt"

// Must have struct with Name string
// AND must implement Greet()
type GreeterConstraint interface {
    ~struct{ Name string } | int   // can use concrete type as type element
    Greet()
	// float64  cannot use primitive type in interface
}

type Person struct {
    Name string
}

func (p Person) Greet() {
    fmt.Println("Hi, I'm", p.Name)
}

type Robot struct {
    Name string
}

func (r Robot) Greet() {
    fmt.Println("Beep, I'm", r.Name)
}

func CallGreet[T GreeterConstraint](g T) {
    g.Greet()
}

func main() {
    CallGreet(Person{"Alice"})
    CallGreet(Robot{"R2D2"})
}
