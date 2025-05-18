package main

import (
	"fmt"
	"slices"
)

func main() {
	// The default value of uninitialized slice is nil
	// var nums []int
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	// var nums = make([]int, 5)
	nums := []int{}
	fmt.Println(len(nums), cap(nums)) // len is the amount of elements in the slice and cap is the maximum capacity of the slice

	nums = append(nums, 5, 6, 7, 8, 9, 10)
	fmt.Println(nums, cap(nums))
	// Note. Capacity doubles as the slice grows

	// copy
	nums2 := make([]int, len(nums))
	copy(nums2, nums)

	// slice operator
	fmt.Println(nums2[0:2])
	fmt.Println(nums2[:2]) // Starts from 0 by default 
	fmt.Println(nums2[0:])	// Ends at the last index by default
	fmt.Println(slices.Delete(nums2, 2, 3))
}
