package main

import "fmt"

func isPalindrome(s string) bool {

	// Convert the string into a slice of runes.
	// This allows us to work with characters directly.
	chars := []rune(s)

	// Create two pointers.
	// left starts at the beginning.
	left := 0

	// right starts at the last character.
	right := len(chars) - 1

	// Continue until the two pointers meet.
	for left < right {

		// Compare the characters at both ends.
		if chars[left] != chars[right] {

			// If they are different, the string
			// cannot be a palindrome.
			return false
		}

		// Move the left pointer toward the middle.
		left++

		// Move the right pointer toward the middle.
		right--
	}

	// If all characters matched, it is a palindrome.
	return true
}

func main() {
	s := "racecar"
	result := isPalindrome(s)
	fmt.Println(result)
}