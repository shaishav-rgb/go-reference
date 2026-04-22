package main

import "fmt"

type Trippler interface {
	Tripple()
}

type Doubler interface {
	Double()
}

type Combined interface {
	Doubler
	Trippler
}

type DoubleInt int

func (d DoubleInt) Double() {
	fmt.Println(d * 2)
}

func (d *DoubleInt) Tripple() {
	fmt.Println((*d) * 3)
}

func DoublerCompare(d1, d2 Doubler) {
	fmt.Println(d1 == d2)
}

func TrippleCompare(d1, d2 Trippler) {
	fmt.Println(d1 == d2)
}

func CombinedCompare(d1, d2 Combined) {
	fmt.Println(d1 == d2)
}
func main() {
	var di DoubleInt = 10
	var di2 DoubleInt = 10
	DoublerCompare(di, di2)
	TrippleCompare(&di, &di2)
	TrippleCompare(di, di2)    // compilation error, value type only has value receiver method. So cannot satisfy the interface Trippler
	CombinedCompare(&di, &di2) // compiles, pointer type has both value receiver and pointer receiver methods.
}
