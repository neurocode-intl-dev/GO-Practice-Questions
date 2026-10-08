package main

import "fmt"


func findMaximum(nums [] int)int {

    max := nums[0]

    for _ , value := range nums {

        if value > max {
            max = value
        }
    }


    return max
}

func main() {

    nums := []int{4, 8, 2, 10}

    result := findMaximum(nums)


    fmt.Println(result)
}