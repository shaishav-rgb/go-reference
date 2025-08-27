package internalchild

import "fmt"


func InternalChildMessage(message string){
	fmt.Printf("This is from internal-child package: %v\n", message)
}