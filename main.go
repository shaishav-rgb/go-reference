package main

import (
	"fmt"
)

// public structs and fields,methods

// type Person struct {
//     Name string
// }

// func (p Person) Greet() {
//     fmt.Println("Hello,", p.Name)
// }

// type Employee struct {
//     Person
//     EmployeeID int
// }

// func main() {
//     e := Employee{
//         Person:     Person{Name: "Alice"},
//         EmployeeID: 123,
//     }

//     e.Greet() // ✅ promoted method from Person
// 	fmt.Println(e.Person)
// }

// type Person struct {
//     Name string
// }

// func (p Person) Greet() {
//     fmt.Println("Hello,", p.Name)
// }

// private structs and fields,methods

type person struct {
    Name string
}

func (p person) greet() {
    fmt.Println("Hello,", p.Name)
}

type employee struct {
    person
    employeeID int
}

func main() {
    e := employee{
        person:     person{Name: "Alice"},
        employeeID: 123,
    }

    e.greet() // ✅ promoted method from person
	fmt.Println(e.person)
}