package main

import "fmt"

// Constraint: must have Name field (string) and Age field (int)
type NamedAged interface {
    ~struct {
        Name string
        Age  int
    }
}

func PrintInfo[T NamedAged](v T) {
    // fmt.Printf("Name: %s, Age: %d\n", v.Name, v.Age)
    fmt.Printf("Name: %s, Age: %d\n")
}

type Person struct {
    Name string
    Age  int
}

type Employee struct {
    Name string
    Age  int
    // ID   int  // Cannot use additional fields, need to have exact match to type terms
}

func main() {
    PrintInfo(Person{"Alice", 30})
    // PrintInfo(Employee{"Bob", 40, 123})
    PrintInfo(Employee{"Bob", 40})
}
