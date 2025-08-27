package parent

import (
	"fmt"
	"test-go-module-package/parent/internal"
	internalchild "test-go-module-package/parent/internal/internal_child"
	// "test-go-module-package/product"
)

func CallInternal(){
	internal.InternalMessage("getting called from parent package")
	internalchild.InternalChildMessage("getting called from parent package")
}

func CircularParentCall(){
	fmt.Println("This is never called because of circular dependency and compilation error")
}

// func CircularCallingProduct(){
// 	product.CircularProductCall() // import cycle, parent is calling product

// }