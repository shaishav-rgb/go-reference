package main

import "fmt"

// Constraint: either int or string
type IntOrString interface {
    int | string
}

// Generic function
func PrintValue[T IntOrString](v T) {
    fmt.Println(v)
}

func PrintSlice[T IntOrString](s []T) {
    for _, v := range s {
        fmt.Println(v)
    }
}

func main() {
    PrintValue(42)       // ✅ int
    PrintValue("hello")  // ✅ string
    // PrintValue(3.14)  // ❌ compile error, float64 not allowed

	PrintSlice([]int{1, 2, 3})
    PrintSlice([]string{"a", "b", "c"})
}

