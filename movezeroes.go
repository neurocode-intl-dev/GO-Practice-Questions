package main

import "fmt"

// moveZeroes moves all zeroes to the end of the slice.
// It maintains the original order of non-zero elements.
func moveZeroes(nums []int) {

    // Create the first pointer.
    // left indicates where the next non-zero number should be placed.
    left := 0

    // Create the second pointer using the for loop.
    // right checks every element in the slice.
    for right := 0; right < len(nums); right++ {

        // Check whether the current number is non-zero.
        if nums[right] != 0 {

            // Swap the current non-zero number with the number at the left pointer.
            nums[left], nums[right] = nums[right], nums[left]

            // Move the left pointer forward.
            // The next non-zero number will be placed here.
            left++
        }
    }

    // No return is required because the original slice is modified directly.
}

func main() {
    nums := []int{0, 1, 0, 3, 12}
    moveZeroes(nums)
    fmt.Println(nums)
}