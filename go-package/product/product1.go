package product

import (
	"fmt"
	"test-go-module-package/payment"
)



func ProductMessage1(message string){
	fmt.Println(message)
	payment.PaymentMessage1(message)

}