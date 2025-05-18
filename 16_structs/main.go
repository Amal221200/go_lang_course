package main

import (
	"fmt"
	"time"
)

type Order struct {
	id        string
	amount    float64
	status    string
	createdAt time.Time // nanosecond precision
}

// Method

// reciever
func (o *Order) changeStatus(status string) { // It should take the reference of the instance to modify the object
	o.status = status
}

func (o Order) getAmount() float64 {
	return o.amount
}

// Constructor
func newOrder(id string, amount float64, status string) *Order {
	return &Order {
		id: id,
		amount: amount,
		status: status,
		createdAt: time.Now(),
	}
}

func main() {
	// If you don't set any field, the property will be set to it's default value
	// order := Order{
	// 	id: "1",
	// 	amount: 34.00,
	// 	status: "recieved",
	// 	createdAt: time.Now(),
	// }

	// order.changeStatus("confirmed")
	// fmt.Println("Order struct", order)
	// fmt.Println("Order amount", order.getAmount())

	// myOrder := newOrder("2", 65, "recieved");
	// fmt.Println(myOrder)

	// Anonymous struct

	language := struct {
		name string
		isGood bool
	}{"golang", true}

	fmt.Println(language)
}