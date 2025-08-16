package main

import "fmt"

// Constraint: either func(int,int)int or func(string,string)string
type AddOrConcat interface {
    func(int, int) int | func(string, string) string
}

func Call[T AddOrConcat](f T, a, b any) any {
    switch fn := any(f).(type) {
    case func(int, int) int:
        return fn(a.(int), b.(int))
    case func(string, string) string:
        return fn(a.(string), b.(string))
    default:
        return nil
    }
}

func execute1() {
    add := func(x, y int) int { return x + y }
    concat := func(x, y string) string { return x + y }

    fmt.Println(Call(add, 2, 3))       // 5
    fmt.Println(Call(concat, "a", "b")) // "ab"
}
