package main

import "fmt"

type dummyStruct struct {
	greetWord string
}

func (d dummyStruct) greeter() string {
	return fmt.Sprintf("Greeting is %v", d.greetWord)
}

type dummy interface {
	greeter() string
}

func returnDummyStruct() dummyStruct {
	return dummyStruct{
		greetWord: "Hello",
	}
}
func returnDummyStructNil() dummyStruct {
	return dummyStruct{
		greetWord: "Hello",
	}
	// return nil; // Cannot use nil as return value for dummyStruct
}
func returnString() string {
	return "Hello"
	// return nil; // Cannot use nil as return value for string
}

func returnSlice() []int {
	// return []int{1,2,3}
	return nil // Can use nil as return value for []int slice
}
func returnMap() map[string]int {
	return map[string]int{}
	// return nil; // Can use nil as return value for map
}

func returnDummyInterfaceValue() dummy {
	return dummyStruct{
		greetWord: "Hello",
	}
}

func returnDummyInterfaceNil() dummy {
	return nil //Can use nil as return value for interface
}

func main() {

}
