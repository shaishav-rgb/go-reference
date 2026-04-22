package main

import "fmt"

type Shape interface {
	Area() float64
}
type Square struct {
	Side float64
}

type Square1 struct {
	Side float64
}

func (s Square1) Area() float64 {
	return s.Side * s.Side
}

func MakeNew[T any]() T {
	var zero T // zero value of type T
	return zero
	
}

func NewCopy[T any](val T) T {
return val
}

func main() {
	sq := MakeNew[Square]() // Square{}
	sq1 := MakeNew[Square1]() // Square1{}
	fmt.Printf("%#v\n", sq)
	fmt.Printf("%#v\n", sq1)

	square2:=Square1{Side: 22}
	square3:=NewCopy(square2)
	fmt.Println(square3)

}