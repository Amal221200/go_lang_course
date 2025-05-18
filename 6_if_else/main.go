package main

import "fmt"

func main() {
	var age int
	fmt.Println("Enter your age")
	fmt.Scanln(&age)

	if age >= 18 {
		fmt.Println("You are an adult", age)
	} else {
		fmt.Println("You are a child", age)
	}
}