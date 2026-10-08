package main

import "fmt"

func removeDuplicates(nums []int) []int {

    seen := map[int] bool{}

    result := []int{}

    for _ , value := range nums {

        if !seen[value] {

            result = append(result, value)
            seen[value] = true
        }
    }

    return result
}

func main() {

    nums := []int{1, 1, 2, 2, 3}

    result := removeDuplicates(nums)

    fmt.Println(result)


	
}