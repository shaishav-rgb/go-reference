package main

import (
	"fmt"
	"sync"
)

// Upgrading to Go 1.22 or later changes the behavior of for loops so they create a new index and value
// variable on each iteration. In Go 1.22  the behavior of a for loop was changed so that it creates new
// variables for the index and value on each iteration instead of reusing a single variable
func main(){
	var wg sync.WaitGroup
	for i:=0;i<10;i++ {
		wg.Add(1)
		//Prior to Go 1.22, this can be solved by either passing a parameter to the goroutine or shadowing and creating a new variable i
		// i:=i
		// go func (i int){
		go func (){
			defer wg.Done()
			fmt.Println(i);
		}()
		// }(i)
	}
	wg.Wait()
}