package main

import (
	"fmt"
)

type Auth struct{}

func (a *Auth) GetUser() string {
	return "Alice"
}

//At first we are only concerned with Struct Auth, no generic struct that has method GetUser() is required at this time

// func TakeAction(a *Auth) {
// 	println("User:", a.GetUser())
// }

// Later we want TakeAction to take any struct which has method GetUser()
type LogicAuth interface {
	GetUser() string
}
func TakeAction(a LogicAuth) {
	println("User:", a.GetUser())
}

func main() {
   authImplementation:=&Auth{}
   fmt.Println(authImplementation.GetUser())
   TakeAction(authImplementation)
}