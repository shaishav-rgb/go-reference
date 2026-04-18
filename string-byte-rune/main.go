package main

import "fmt"

//Also Reference pg-54, learning go-joe bodner
func main() {
	var s string = "Hello there"
	// prints bytes
	fmt.Println(s[0])
	// print character after conversion from byte to string
	fmt.Println(string(s[0]))
	//print characters
	fmt.Println(s[0:5])
	var b byte = s[6]
	fmt.Println(b)
}


