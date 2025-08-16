package main

import (
	"errors"
	"fmt"
)

// func divAndRemainder(num, denom int) (int, int, error) {
func divAndRemainder1(num, denom int) (float64, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return float64(num / denom), num % denom, nil
}

func main() {
	// division,remainder,_:=divAndRemainder(5,2);
	division, remainder, _ := divAndRemainder1(5, 2)
	fmt.Printf("Division:%f\nRemainder:%d\n", division, remainder)

}
