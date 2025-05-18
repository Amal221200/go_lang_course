package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	paymentGw paymenter
}

func (p payment) makePayment(amount float32) {
	p.paymentGw.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("Making payment of", amount, "from razorpay")
}

type stripe struct{}

func (r stripe) pay(amount float32) {
	fmt.Println("Making payment of", amount, "from stripe")
}

type paypal struct{}

func (r paypal) pay(amount float32) {
	fmt.Println("Making payment of", amount, "from paypal")
}

func main() {
	// gateway := razorpay{}
	// gateway := stripe{}
	gateway := paypal{}

	newPayment := payment {
		paymentGw: gateway,
	}
	newPayment.makePayment(5600)
}
