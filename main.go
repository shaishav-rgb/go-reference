package main

import (
	"fmt"
)

type IntChannel interface {
	chan<- int | <-chan int
}

func useChannel[C IntChannel](ch C) {
	fmt.Println("Got a channel:", ch)
}
func main() {

	sendCh := make(chan int)
	recvCh := make(chan int)

	var _ = (chan<- int)(sendCh) // can use both variable as _
	var _ = (<-chan int)(recvCh) // can name both variable as _
	// fmt.Println(_) // cannot use _ as value or type

	var sendChannel = (chan<- int)(sendCh)
	var receiveChannel = (<-chan int)(recvCh)

	useChannel(sendChannel)
	// useChannel(sendCh)  // cannot use sendCh here, have to do type conversion
	useChannel(receiveChannel)

	//cannot use type terms as variable type, can be only used as generic type constraint
	// var _ IntChannel = (chan<- int)(sendCh)
	// var _ IntChannel = (<-chan int)(recvCh)

}
