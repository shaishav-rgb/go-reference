package product

import (
	"fmt"
	// . "test-go-module-package/payment"  //using dot import
	// "test-go-module-package/parent/internal"       // compilation error, use of internal package not allowed
	"test-go-module-package/parent"
	// renamePayment "test-go-module-package/payment" // overriding a package name when importing
	// . "test-go-module-package/payment" //using dot import
	_ "test-go-module-package/init-package"
)

// func PaymentMessage(message string){ // Payment Message already declared through dot import so cannot re-declare it again here, compilation error
// 	fmt.Println(message)
// }

// var renamePayment=3   //renamePayment already declared through package variable override, cannot have variable of the same name inside the same block(package block, they say file block gets promoted to package block)


func ProductMessage(message string){
	fmt.Println(message)
	// renamePayment.PaymentMessage(message)
	// renamePayment:=3    // renamePayment can be shadowed here inside a sub-block
	// fmt.Println(renamePayment)
}

// func callFromProduct(){
// 	internal.InternalMessage("Error")  // internal not allowed!!!
// 	internal.CallInternal()
// }


func CircularProductCall(){
	fmt.Printf("This is never called because of circular dependency and compilation error\n")
	parent.CallInternal() // import cycle here, product is calling
}