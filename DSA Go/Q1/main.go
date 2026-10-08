package main

import "fmt"

func twoSum(nums []int, target int) []int {

    numbers := map[int]int{}

    for i, value := range nums {

        needed := target - value

        index, ok := numbers[needed]

        if ok {
            return []int{index, i}
        }

        numbers[value] = i
    }

    return []int{}


}

func main() {

    nums := []int{2, 7, 11, 15}
    target := 9

    result := twoSum(nums, target)
    fmt.Println(result)
}

//[0 1]