package main

import "fmt"

func main() {
	var nums [4]int
	nums[0] = 1
	nums[1] = 2
	nums[2] = 3
	nums[3] = 4

	// Another way to initialize an array
	// var nums = [4]int{1, 2, 3, 4}

	fmt.Println(len(nums), nums)

	for i, val := range nums {
		fmt.Println(i, val)
	}

	// 2d arrays
	matrix := [2][2]int{{1, 2}, {1, 2}}
	fmt.Println(matrix)
}
