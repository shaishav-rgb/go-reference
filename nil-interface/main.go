package main

import "fmt"

type Greeter interface {
	Greet()
}
type Person struct {
	Name string
}

// func (p *Person) Greets() { //cannot be Greets, must match interface Greeter
func (p *Person) Greet() {
	if p == nil {
		fmt.Println(p == nil)
		fmt.Printf("%v is nil\n", p)
		return
	}
	fmt.Println(p)
	fmt.Println("Hello,", p.Name) //panics if p is nil
}
func main() {
	var p *Person
	var g Greeter = p     // type = *Person, value = nil
	fmt.Println(g == nil) // false (because type is set)
	fmt.Printf("%v is nil\n", g)
	g.Greet()
}
