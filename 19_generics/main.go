package main

import "fmt"

// func printSlice[T int | string](items []T) {
// 	for _, v := range items {
// 		fmt.Println(v)
// 	}
// }
func printSlice[T comparable](items []T) {
	for _, v := range items {
		fmt.Println(v)
	}
}

type Stack[T int | string] struct {
	elements []T
}

func main() {
	printSlice([]int{1, 2, 3})
	printSlice([]string{"Amal", "Ashin"})
}
