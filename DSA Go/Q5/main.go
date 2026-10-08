package main

import "fmt"

func isPalindrome(word string) bool {

    left := 0
    right := len(word) - 1

    for left < right {

        if word[left] != word[right] {
            return false
        }

        left++
        right--
    }

    return true
}

func main() {

    text := "racecar"
    result := isPalindrome(text)

    fmt.Println(result)
}

//true