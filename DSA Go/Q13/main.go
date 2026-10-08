package main

import "fmt"

func rotate(nums []int, k int) {

	k = k % len(nums)

	for i := 0; i < k; i++ {

		last := nums[len(nums)-1]

		for j := len(nums) - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}

		nums[0] = last
	}
}

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7}

	rotate(nums, 3)

	fmt.Println(nums)
}