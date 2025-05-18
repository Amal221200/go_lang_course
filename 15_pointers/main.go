package main

import "fmt"

func main() {
	num := 90

	changeNum(&num)

	fmt.Println("After changeNum in main", num)
}

// by value
// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("In changeNum", num)
// }

// by address
func changeNum(num *int) {
	*num = 5
	fmt.Println("In changeNum", *num)
}
