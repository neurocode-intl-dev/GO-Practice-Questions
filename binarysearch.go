package main

import "fmt"

// binarySearch searches for the target value in a sorted slice.
// It returns the index of the target if found.
// If the target is not found, it returns -1.
func binarySearch(nums []int, target int) int {

	// left points to the beginning of the search area.
	left := 0

	// right points to the end of the search area.
	right := len(nums) - 1

	// Continue searching while the search area is valid.
	for left <= right {

		// Find the middle index.
		mid := left + (right-left)/2

		// Check if the middle element is the target.
		if nums[mid] == target {

			// Target found.
			// Return its index.
			return mid
		}

		// If the middle element is smaller than the target, the target must be on the right side.
		if nums[mid] < target {

			// Ignore the left half. Start searching from the element after the middle.
			left = mid + 1

		} else {

			// The middle element is greater than the target.
			// Therefore, search the left half.
			right = mid - 1
		}
	}

	// Target was not found in the array.
	return -1
}

func main() {
	nums := []int{1, 3, 5, 7, 9}
	target := 7
	result := binarySearch(nums, target)
	fmt.Println(result)
}