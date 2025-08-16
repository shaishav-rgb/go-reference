package main

import "fmt"

// Define a type element interface
type StringOrIntSlice interface {
	// ~string | ~[]int
	string | []int
}

// Generic function using the type element
func PrintValue[T StringOrIntSlice](val T) {
	switch any(val).(type) {
	case string:
		fmt.Println("String:", val)
	case []int:
		fmt.Println("Int slice:", val)
	}
}

type Container[T StringOrIntSlice] struct {
	Value T
}

func main() {
	PrintValue("Hello")                          // String             // type parameter is optional in function
	PrintValue([]int{1, 2, 3})                   // Int slice
	c1 := Container[string]{Value: "Hello"}      // type parameter is mandatory in struct generic
	c2 := Container[[]int]{Value: []int{10, 20}} // type parameter is mandatory in struct generic
	fmt.Println(c1.Value)
	fmt.Println(c2.Value)

	type userIntSlice []int
	userIntSliceValues := userIntSlice{10, 20}               // can use []int as a user-defined type and do this
	// c3 := Container[userIntSlice]{Value: userIntSliceValues} // if omit ~ in type element then this is error
	// fmt.Println(c3.Value)
	c4 := Container[[]int]{Value: userIntSliceValues} // can use "userIntSliceValues" as  []int as userIntSlice is derived from []int
	fmt.Println(c4.Value)
}
