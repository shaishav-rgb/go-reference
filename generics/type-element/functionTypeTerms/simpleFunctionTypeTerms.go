package main

import "fmt"

// Constraint: any function that takes two ints and returns an int
type IntBinOp interface {
    func(int, int) int
}

// Generic function that accepts a function
func Apply[T IntBinOp](f T, a, b int) int {
    return f(a, b)
}

func main() {
    add := func(x, y int) int { return x + y }
    mul := func(x, y int) int { return x * y }

    fmt.Println(Apply(add, 2, 3)) // 5
    fmt.Println(Apply(mul, 2, 3)) // 6

    // fmt.Println(Apply(func(x string, y string) string { return x + y }, "a", "b"))
    // ❌ Compile error: signature does not match IntBinOp
}