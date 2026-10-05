package main

import "fmt"

// rotate rotates the array to the right by k positions.
func rotate(nums []int, k int) {

	// If the array is empty, there is nothing to rotate.
	if len(nums) == 0 {
		return
	}

	// If k is greater than the length of the array,
	// we only need the remainder.
	//
	// Example:
	// length = 7
	// k = 10
	//
	// 10 % 7 = 3
	//
	// Rotating 10 times is the same as rotating 3 times.
	k = k % len(nums)

	// Create a new slice to store the rotated array.
	result := make([]int, len(nums))

	// Loop through every element in the original array.
	for i := 0; i < len(nums); i++ {

		// Calculate the new position of the current element.
		//
		// (i + k) % len(nums)
		//
		// This makes the elements that go past
		// the end of the array wrap around to the beginning.
		newIndex := (i + k) % len(nums)

		// Put the current number into its new position.
		result[newIndex] = nums[i]
	}

	// Copy the rotated values back into the original slice.
	copy(nums, result)
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	fmt.Println(nums)
}