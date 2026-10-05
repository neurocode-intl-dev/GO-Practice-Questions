package main

import "fmt"

// findMaximum returns the largest number in the slice.
func findMaximum(nums []int) int {

	// Assume the first element is the maximum initially.
	max := nums[0]

	// Loop through the slice starting from the second element.
	for i := 1; i < len(nums); i++ {

		// Check if the current number is greater than the current maximum.
		if nums[i] > max {

			// If it is greater, update max.
			max = nums[i]
		}
	}

	// Return the largest number.
	return max
}

func main() {
	nums := []int{4, 8, 2, 10}
	result := findMaximum(nums)
	fmt.Println(result)
}