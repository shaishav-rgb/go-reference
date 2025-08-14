package main

import (
	"fmt"
)

func LogOutput(message string) {
	fmt.Println(message)
}
func LogOutput1(message int) {
	fmt.Println(message)
}

type Logger interface {
	Log(message string)
}

type LoggerAdapter func(message string)

func (lg LoggerAdapter) Log(message string) {
	lg(message)
}

func main() {
	normalFunc := LogOutput
	normalFunc("Hello")

	typeConvertedFunc := LoggerAdapter(LogOutput)
	typeConvertedFunc("Bye")
	// castedFunc1:=LoggerAdapter(LogOutput1)  //cannot convert LogOutput1 (value of type func(message int)) to type
	// castedFunc1("Bye");
}
