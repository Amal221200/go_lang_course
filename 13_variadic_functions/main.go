package main

import "fmt"

// Functions that can take n number of arguments are called variadic functions, eg: fmt.Println(1, 2, 3, "hello")
func main() {
	// fmt.Println(average(23, 45, 27))
	nums := []float64{ 33, 44, 55}
	fmt.Println(average(nums...)) // Spread syntax
}

func average(nums ...float64) float64 { // Rest syntax
	total := 0.0

	for _, val := range nums {
		total += val
	}

	return total / float64(len(nums))
}
