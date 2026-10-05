package main

import "fmt"

// maxSubArray finds the largest sum of a contiguous subarray.
func maxSubArray(nums []int) int {

	// Start with the first number as the current sum.
	// currentSum represents the maximum sum of a subarray ending at the current position.
	currentSum := nums[0]

	// maxSum stores the largest sum we have found so far.
	maxSum := nums[0]

	// Start from the second element.
	for i := 1; i < len(nums); i++ {

		// Decide whether to:
		// 1. Start a new subarray from nums[i]
		// OR
		// 2. Add nums[i] to the existing subarray.
		// If currentSum + nums[i] is smaller than nums[i], it is better to start a new subarray.
		if currentSum+nums[i] < nums[i] {

			// Start a new subarray.
			currentSum = nums[i]

		} else {

			// Continue the existing subarray.
			currentSum += nums[i]
		}

		// If the current sum is greater than the maximum sum found so far, update maxSum.
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	// Return the largest subarray sum.
	return maxSum
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	result := maxSubArray(nums)
	fmt.Println(result)
}