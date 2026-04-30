package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Item string

type Order struct {
	ID          string    `json:"id"`
	Items       []Item    `json:"items"`
	DateOrdered time.Time `json:"date_ordered"`
	CustomerID  string    `json:"customer_id"`
}

func (o Order) MarshalJSON() ([]byte, error) {
	type Dup Order
	tmp := struct {
		DateOrdered string `json:"date_ordered"`
		Dup
	}{
		Dup: Dup(o),
		// Dup: (Dup)(o), // no need for small brackets here
	}
	tmp.DateOrdered = o.DateOrdered.Format(time.RFC822Z)
	b, err := json.Marshal(tmp)
	fmt.Println("Inside MarshalJSON")
	return b, err

}

// Explicit path — through the embedded type's name
// tmp.Dup.ID
// tmp.Dup.CustomerID
// tmp.Dup.Items
// tmp.Dup.DateOrdered

// Promoted path — as if the fields were directly on tmp
// tmp.ID
// tmp.CustomerID
// tmp.Items
// tmp.DateOrdered  ← AMBIGUOUS!

// tmp.DateOrdered       // refers to the OUTER string field
// tmp.Dup.DateOrdered   // refers to the INNER time.Time field

func (o *Order) UnmarshalJSON(b []byte) error {
	type Dup Order
	tmp := struct {
		DateOrdered string `json:"date_ordered"`
		*Dup
	}{
		Dup: (*Dup)(o),
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}
	o.DateOrdered, err = time.Parse(time.RFC822Z, tmp.DateOrdered)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	type Order1 Order
	o := Order{
    ID:          "ord-42",
    CustomerID:  "cust-7",
    Items:       []Item{"a","b"},
    DateOrdered: time.Date(2026, 4, 29, 15, 4, 5, 0, time.UTC),
}
	o1 := Order1{
    ID:          "ord-42",
    CustomerID:  "cust-7",
    Items:       []Item{"a","b"},
    DateOrdered: time.Date(2026, 4, 29, 15, 4, 5, 0, time.UTC),
}

output,_:=json.Marshal(o)
fmt.Println(string(output))
output1,_:=json.Marshal(o1)
fmt.Println(string(output1))
}
