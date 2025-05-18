package main

import (
	"fmt"
	"time"
)

func main() {
	var choice string
	fmt.Println("Select an option: A, B, C")
	fmt.Scanln(&choice)

	// simple switch
	switch choice {
	case "A":
		fmt.Println("Option A")
	case "B":
		fmt.Println("Option B")
	case "C":
		fmt.Println("Option C")
	case "":
		fmt.Println("No options selected")
	default:
		fmt.Println("Your option is not available")
	}

	// multiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}

	// type switch
	whoAmI := func(i any) {
		switch t := i.(type) {
		case int:
			fmt.Println("It's an integer")
		case string:
			fmt.Println("It's a string")
		case float64:
			fmt.Println("It's a float")
		default:
			fmt.Println("Other", t)
		}

	}

	whoAmI(34)
}
