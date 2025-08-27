package internal

import (
	"fmt"
	internalchild "test-go-module-package/parent/internal/internal_child"
)


func InternalMessage(message string){
	fmt.Printf("This is from internal package: %v\n", message)
}

func CallInternal(){
	internalchild.InternalChildMessage("getting called from parent internal package")
}