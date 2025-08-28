package main //package should be named main to run
// package main1 //Error, package should be named main to run

import (
	"fmt"
	// initpackage "test-go-module-package/init-package" // referencing initpackage so init gets called before anything in main function

	// "test-go-module-package/parent/sibling"
	// "test-go-module-package/product"
	"github.com/spf13/cobra"
)

// func main1(){      //Error, function should be named main to run
func main() { //function should be named main to run
	// fmt.Println("Calling main function")
	// initpackage.SayHi()
	// product.ProductMessage("This is product message")
	// parent.CallInternal()
	// sibling.CallInternal()
	// product.CircularProductCall()
	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, Modules!")

		},
	}
	fmt.Println("Calling cmd.Execute()!")
	cmd.Execute()
}
