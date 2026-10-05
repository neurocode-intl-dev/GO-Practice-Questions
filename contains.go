package main

import "fmt"

// containsDuplicate checks whether the slice contains any duplicate value.
func containsDuplicate(nums []int) bool {

	// Create a hash map to work as a Set.
	// The key is the number from the slice.
	// The value tells us whether we have already seen the number.
	seen := make(map[int]bool)

	// Loop through every number in the slice.
	// _ means we don't need the index.
	// num contains the current number.
	for _, num := range nums {

		// Check whether the current number already exists in the Set.
		if seen[num] {

			// If the number already exists, we found a duplicate.
			return true
		}

		// The number is not a duplicate, so add it to the Set.
		seen[num] = true
	}

	// If the loop finishes without finding a duplicate, return false.
	return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	result := containsDuplicate(nums)
	fmt.Println(result)
}