package main

import "fmt"

func main() {

	// 1. while loop implementation
	// i := 0
	// for i < 1 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// 2. infinite loop implementation
	// for {
	// 	fmt.Println(i)
	// }

	// 3. normal for loop
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// 4. range
	for i := range 10 {
		fmt.Println(i)
	}
}