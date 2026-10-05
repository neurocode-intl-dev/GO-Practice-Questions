package main

import "fmt"

// twoSum finds two numbers in the slice whose sum equals the target.
// It returns the indexes of those two numbers.
func twoSum(nums []int, target int) []int {    // nums []int -> slice of integers
	// create the map
    seen := make(map[int]int)

	//i → index, num → value
    for i, num := range nums {
		//The complement is the number we need to find to reach the target
        complement := target - num

		// Check whether the complement already exists in the map.
        // j      -> index where the complement was found
        // exists -> true if the complement exists in the map
        if j, exists := seen[complement]; exists {
            return []int{j, i}
        }

		// This stores the current number and its index in the map.
        seen[num] = i
    }

	// If no two numbers are found, return an empty slice.
    return []int{}
}

func main() {
    nums := []int{2, 7, 11, 15}
    target := 9

    result := twoSum(nums, target)

    fmt.Println(result)
}