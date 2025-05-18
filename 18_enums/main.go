package main

import "fmt"

// type OrderStatus int

// const (
// 	Recieved OrderStatus = iota
// 	Confirmed
// 	Prepared
// 	Delivered
// )
type OrderStatus string

const (
	Recieved  OrderStatus = "recieved"
	Confirmed  = "confirmed"
	Prepared   = "prepared"
	Delivered  = "delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing status to", status)
}

func main() {
	changeOrderStatus(Delivered)
}
