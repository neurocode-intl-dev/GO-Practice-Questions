package main

import "fmt"

func binarySearch(nums []int, target int) int {

	left := 0
	right := len(nums) - 1

	for left <= right {

		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {

	nums := []int{1, 3, 5, 7, 9}

	result := binarySearch(nums, 7)

	fmt.Println(result)
}