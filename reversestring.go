package main

import "fmt"

func reverseString(s string) string {
	// Convert string to []rune
	chars := []rune(s)

	// Two pointers
	left := 0
	right := len(chars) - 1

	// Move pointers toward the middle
	for left < right {
		// Swap characters
		chars[left], chars[right] = chars[right], chars[left]

		// Move left pointer forward
		left++

		// Move right pointer backward
		right--
	}

	return string(chars)
}

func main() {
	s := "hello"

	result := reverseString(s)

	fmt.Println(result)
}