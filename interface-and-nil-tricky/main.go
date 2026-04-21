package main

import "fmt"



type MyError struct{}

func (e *MyError) Error() string {
    return "something went wrong"
}
type MyStructTypeConcrete struct{
	name string
}

func (e *MyStructTypeConcrete) Hey() string {
    return "something went wrong"
}

type MyStructType interface{
	Hey() string
}

func doWork() error {
    var err *MyError
  

	fmt.Println(err==nil) // err is nil

	var numberPointer *int
	fmt.Println(numberPointer==nil) //numberPointer is nil

    return err  // This is where the transformation happens. The function's return type is error, which is an interface. The variable err is *MyError, a concrete pointer type. When you return a concrete value from a function whose return type is an interface, Go wraps the concrete value into an interface value.          
}

func doWork1() MyStructType {
    var myStruct *MyStructTypeConcrete
  

	fmt.Println(myStruct==nil) // myStruct is nil

	var numberPointer *int
	fmt.Println(numberPointer==nil) //numberPointer is nil

    return myStruct  // This is where the transformation happens. The function's return type is MyStructType, which is an interface. The variable myStruct is *MyStructTypeConcrete, a concrete pointer type. When you return a concrete value from a function whose return type is an interface, Go wraps the concrete value into an interface value.          
}







func main() {
    err := doWork1()
	fmt.Println(err) //prints <nil> but err==nil is false
	fmt.Println(err==nil) //always false as err is interface
    if err != nil {
        fmt.Println("got an error!")  // ← this prints!
    } else {
        fmt.Println("no error")
    }
}

// reference, the above is similar to this(folder "nil interface" in this repo) 
// package main

// import "fmt"

// type Greeter interface {
// 	Greet()
// }
// type Person struct {
// 	Name string
// }

// // func (p *Person) Greets() { //cannot be Greets, must match interface Greeter
// func (p *Person) Greet() {
// 	if p == nil {
// 		fmt.Println(p == nil)
// 		fmt.Printf("%v is nil\n", p)
// 		return
// 	}
// 	fmt.Println(p)
// 	fmt.Println("Hello,", p.Name) //panics if p is nil
// }
// func main() {
// 	var p *Person
// 	var g Greeter = p     // type = *Person, value = nil
// 	fmt.Println(g == nil) // false (because type is set)
// 	fmt.Printf("%v is nil\n", g)
// 	g.Greet()
// }