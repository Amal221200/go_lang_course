package main

import "fmt"

func main() {
	const PI = 3.14

	// Multi-line constants and variables declaration

	const (
		port = 5000
		host = "localhost"
	)

	var (
		name = "Amal"
		age  = 45
	)
	fmt.Println(name, age, port, host)
}