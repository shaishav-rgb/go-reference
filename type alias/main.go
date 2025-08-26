package main

import "fmt"

type Foo struct {
	x int
	S string
}

func (f Foo) Hello() string {
	return "hello"
}
func (f Foo) goodbye() string {
	return "goodbye"
}

type Bar = Foo

func (f Bar) Hello1() string {  //this is same as making method on Foo struct
	return "hello1"
}

func main() {
	MakeBar()
}

func MakeBar() Bar {
	bar := Bar{
		x: 20,
		S: "Hello",
	}
	var f Foo = bar
	fmt.Println(f.Hello())
	fmt.Println(f.Hello1())   // method can be used on both original and aliased type
	fmt.Println(bar.Hello1())   // method can be used on both original and aliased type
	return bar
}
