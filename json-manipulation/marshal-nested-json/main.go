package main

import (
	"encoding/json"
	"fmt"
)

type Customer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Item struct {
	SKU string `json:"sku"`
	Qty int    `json:"qty"`
}

// type Order struct {
//     ID       string   `json:"id"`
//     Customer Customer `json:"customer"`
//     Items    []Item   `json:"items"`
// }

type Order struct {             //verbose struct type
	ID       string `json:"id"`
	Customer struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"customer"`
	Items []struct {
		SKU string `json:"sku"`
		Qty int    `json:"qty"`
	} `json:"items"`
}

func main() {
	// o := Order{
	// 	ID:       "ord-42",
	// 	Customer: Customer{Name: "Alice", Email: "alice@example.com"},
	// 	Items: []Item{
	// 		{SKU: "A1", Qty: 2},
	// 		{SKU: "B7", Qty: 1},
	// 	},
	// }
	o := Order{   //initializing verbose struct
		ID: "ord-42",
		Customer: struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}{Name: "Alice", Email: "alice@example.com"},
		Items: []struct {
			SKU string `json:"sku"`
			Qty int    `json:"qty"`
		}{
			{SKU: "A1", Qty: 2},
			{SKU: "B7", Qty: 1},
		},
	}
	b, _ := json.Marshal(o)
	fmt.Println(string(b))
}
