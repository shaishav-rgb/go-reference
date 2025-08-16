package main

import "fmt"

// Constraint: function taking int, returning T
type IntTo[T any] interface {
    func(int) T
}

// Generic function applying f to a value
func ApplyFunc[T any, F IntTo[T]](f F, x int) T {
    return f(x)
}

func execute() {
    square := func(x int) int { return x * x }
    greet := func(x int) string { return fmt.Sprintf("Hello #%d", x) }

    fmt.Println(ApplyFunc(square, 5)) // 25
    fmt.Println(ApplyFunc(greet, 3))  // "Hello #3"
}
