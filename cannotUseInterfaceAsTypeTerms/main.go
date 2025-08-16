package main

import "fmt"


type Speaker interface {
    func(string) string | Speakable   //Cannot use interface as type terms. Only concrete types are allowed
}

type Speakable interface {
    Speak() string
}

func Call[T Speaker](s T) {
    switch v := any(s).(type) {
    case func(string) string:       // using function type in type switch
        fmt.Println(v("World"))
    case Speakable:
        fmt.Println(v.Speak())
    }
}

type Person struct {
    Name string
}

func (p Person) Speak() string {
    return "Hello, I'm " + p.Name
}

func main() {
    f := func(s string) string { return "Hi " + s }
    Call(f)             // uses function
    Call(Person{"Bob"}) // uses struct+method
}
