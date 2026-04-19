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
	//the type of b is byte
	 b:= s[6]
	fmt.Println(b)

	fmt.Println("------------------------------")

	var s1 string = "Hello😎"

	b1:=[]byte(s1)
	r1:=[]rune(s1)

	for _,value:=range b1{
		fmt.Println(value)
		fmt.Println(string(value))
	}
	fmt.Println("------------------------------")
	
	for _,value:=range r1{
		fmt.Println(value)
		fmt.Println(string(value))
	}

	fmt.Println("------------------------------")

	//type of value is rune
	for _,value:=range s1{
		fmt.Println(value)
		fmt.Println(string(value))
	}
}


