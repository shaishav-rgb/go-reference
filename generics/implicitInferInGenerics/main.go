package main

import (
	"fmt"
	"strconv"
)

type Integer interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

func Convert[T1, T2 Integer](in T1) T2 {
	return T2(in)
}

func Convert1[T1 ~int, T2 ~string](in T1, s T2) T2 {
// func Convert1[T1 ~int, T2 string](in T1, s T2) T2 {
	newString:=strconv.Itoa(int(in))+"s"
	return T2(newString)
}

// func Convert1[T1 ~int, T2 string](in T1, s T2) T2 {
// 	newString:=strconv.Itoa(int(in))+"s"
// 	return T2(newString)
// }
func main() { 
	var a int = 10
	type newString string
	 var stringValue newString= "hi"
	b1 := Convert1(a,stringValue) // inferred as newString type. Parameter is also a user-defined type built from built-in type, can accept this  user-defined type because of "~" in front of type in parameter.
	fmt.Println(b1)
	b := Convert[int, int64](a) // can't infer the return type, in64 is not used inside the function, it is only a generic type.
	fmt.Println(b)
}
