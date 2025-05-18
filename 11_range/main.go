package main

import "fmt"

func main() {
	nums := []int{6, 7, 8}

	// Traditional way
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	// Range
	for index, num := range nums {
		fmt.Println(index, num)
	}

	// Range with map
	m := map[string]string{"name": "Amal", "color": "blue"}

	for key, val := range m {
		fmt.Println(key, val)
	}

	// Range with string
	name := "Amal"

	// Characters are called rune:- ex 'A', 'z'
	// The first variable (ind) will store the starting point from where the rune is taking space.
	for ind, char := range name {
		// String is a collection of runes, i.e if we print the rune you will see the unicode/ascii value of that character
		// fmt.Println(ind, char)

		// To convert a rune to string, use string() function.
		fmt.Println(ind, string(char))
	}

}
