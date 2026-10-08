package main

import "fmt"

func missingNumber(nums []int) int {

	n := len(nums)

	expectedSum := n * (n + 1) / 2

	actualSum := 0

	for _ , value := range nums {
		actualSum = actualSum + value
	}

	return expectedSum - actualSum
}

func main() {

	nums := []int{3, 0, 1}

	result := missingNumber(nums)

	fmt.Println(result)
}