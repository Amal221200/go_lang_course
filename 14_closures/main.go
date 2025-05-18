package main

import "fmt"

func Counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	counter := Counter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}