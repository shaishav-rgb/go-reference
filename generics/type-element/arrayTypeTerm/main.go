package main

import "fmt"

// Constraint: must be [3]int (array of exactly 3 ints)
type ThreeInts interface {
    ~[3]int
}

func SumArray[T ThreeInts](a T) int {
    sum := 0
    for _, v := range a {
        sum += v
    }
    return sum
}

func main() {
    arr := [3]int{1, 2, 3}
    fmt.Println(SumArray(arr)) // ✅ Works: 6

    // arr2 := [4]int{1, 2, 3, 4}
    // fmt.Println(SumArray(arr2)) // ❌ compile error: wrong length
}