package main

import (
	"encoding/json"
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
	return b, err

}

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

func main() {}
