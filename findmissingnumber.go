package main

import "fmt"

// findMissingNumber finds the missing number from a slice containing numbers from 0 to n.
func findMissingNumber(nums []int) int {

	// The length of the slice is n.
	//
	// Example:
	// nums = [3, 0, 1]
	// len(nums) = 3
	//
	// Therefore, numbers should be from 0 to 3.
	n := len(nums)

	// Calculate the expected sum of numbers from 0 to n.
	//
	// Formula:
	// n * (n + 1) / 2
	//
	// Example:
	// 3 * (3 + 1) / 2
	// = 6
	expectedSum := n * (n + 1) / 2

	// Create a variable to store the actual sum of the numbers in the slice.
	actualSum := 0

	// Loop through every number in the slice.
	for _, num := range nums {

		// Add the current number to actualSum.
		actualSum += num
	}

	// The difference between the expected sum and actual sum is the missing number.
	//
	// Example:
	// expectedSum = 6
	// actualSum   = 4
	//
	// 6 - 4 = 2
	return expectedSum - actualSum
}

func main() {
	nums := []int{3, 0, 1}
	result := findMissingNumber(nums)
	fmt.Println(result)
}