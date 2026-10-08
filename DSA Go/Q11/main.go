package main

import "fmt"

func moveZeroes(nums []int) {

	position := 0

	for _, value := range nums {

		if value != 0 {
			nums[position] = value
			position++
		}
	}

	for position < len(nums) {
		nums[position] = 0
		position++
	}
}

func main() {

	nums := []int{0, 1, 0, 3, 12, 0}
	moveZeroes(nums)

	fmt.Println(nums)
}