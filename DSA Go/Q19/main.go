package main

import "fmt"

func maxSubArray(nums []int) int {


	currentSum := nums[0]

	maxSum := nums[0]

	for i := 1; i < len(nums); i++ {

		if currentSum + nums[i] < nums[i] {
			currentSum = nums[i]
		} else {
			currentSum = currentSum + nums[i]

		}
		fmt.Println(currentSum)
		if currentSum > maxSum {
			maxSum = currentSum
		}
		fmt.Println(maxSum)
	}


	return maxSum
}

func main() {
	
	nums := []int{-2, 1, -3, 4, -1, 2, -5, 4}

	result := maxSubArray(nums)
	
	fmt.Println(result)
}

//5