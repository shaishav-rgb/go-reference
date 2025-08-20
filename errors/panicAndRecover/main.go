package main

import "fmt"

func doPanic(msg string) {
	panic(msg)
}

func innerFunc(msg string) {
	defer func() {							//defer recovery should be registered before panic happens
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()
	doPanic(msg)
}

func main() {
	// defer func() {							//defer recovery should be registered before panic happens
	// 	if v := recover(); v != nil {
	// 		fmt.Println(v)
	// 	}
	// }()
	innerFunc("Hello")

}
