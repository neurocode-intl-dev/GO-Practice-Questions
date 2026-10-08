package main

import "fmt"

func containsDuplicate(nums []int) bool {



    numbers := map[int] bool{}

    for _ , value := range nums{

        if numbers[value] {
            return true
        }
        numbers[value] = true
    }

    return false
}



func main() {

    nums := []int{1,2,3,1,4,4}
    result := containsDuplicate(nums)

    fmt.Println(result)
}

//true