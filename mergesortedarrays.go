package main

import "fmt"

// mergeSortedArrays merges two sorted arrays into one sorted array.
func mergeSortedArrays(nums1 []int, nums2 []int) []int {

	// Create an empty slice to store the merged sorted values.
	result := []int{}

	// i is the pointer for the first array.
	i := 0

	// j is the pointer for the second array.
	j := 0

	// Continue while both pointers are inside their respective arrays.
	for i < len(nums1) && j < len(nums2) {

		// Compare the current elements of both arrays.
		if nums1[i] < nums2[j] {

			// nums1[i] is smaller, so add it to the result.
			result = append(result, nums1[i])

			// Move the first pointer forward.
			i++

		} else {

			// nums2[j] is smaller or equal, so add it to the result.
			result = append(result, nums2[j])

			// Move the second pointer forward.
			j++
		}
	}

	// If there are remaining elements in the first array, add them to the result.
	for i < len(nums1) {
		result = append(result, nums1[i])

		// Move the first pointer forward.
		i++
	}

	// If there are remaining elements in the second array, add them to the result.
	for j < len(nums2) {
		result = append(result, nums2[j])

		// Move the second pointer forward.
		j++
	}

	// Return the final sorted array.
	return result
}

func main() {
	nums1 := []int{1, 3, 5}
	nums2 := []int{2, 4, 6}
	result := mergeSortedArrays(nums1, nums2)
	fmt.Println(result)
}