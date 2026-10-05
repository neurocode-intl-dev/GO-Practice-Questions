package main

import "fmt"

// removeDuplicates returns a new slice containing only unique values.
func removeDuplicates(nums []int) []int {

	// Create a map to work as a Set.
	// The key stores the number.
	// The value tells us whether we have already seen it.
	seen := make(map[int]bool)

	// Create an empty slice to store unique numbers.
	result := []int{}

	// Loop through every number in the input slice.
	for _, num := range nums {

		// Check whether the number already exists in our Set.
		if seen[num] {
			// Number is already present, so skip it.
			continue
		}

		// Add the number to the Set.
		seen[num] = true

		// Add the unique number to the result slice.
		result = append(result, num)
	}

	// Return the slice containing only unique values.
	return result
}

func main() {
	nums := []int{1, 1, 2, 2, 3}
	result := removeDuplicates(nums)
	fmt.Println(result)
}