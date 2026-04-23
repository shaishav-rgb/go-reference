package main

import "fmt"

type Shape interface {
	Area() float64
}

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func CompareShapes(a, b Shape) bool {
	return a == b

}

type ComparableShape interface {
	comparable
	Area() float64
}

func CompareShapesGeneric[T ComparableShape](a,b T) bool{
	return a == b
}

func main() {
	square := Square{Side: 4}
	rectangle := Rectangle{Width: 4, Height: 2}

	isEqual:=CompareShapes(square, rectangle) // in this non-generic version, the comparison result is false but no compilation check. If concrete type is non-comparable like slice then panics at runtime
	fmt.Println(isEqual)
	
	isEqual1:=CompareShapesGeneric(square, rectangle) // compilation error, in this generic version concrete type should also match.
	fmt.Println(isEqual1)

	var square1 ComparableShape=Square{Side:4} // compilation error, comparable in ComparableShape is a type constraint so cannot be used like this
	// var square1 Shape=Square{Side:4} // no comparable in interface Shape so can be used like this
	fmt.Println(square1)
}
