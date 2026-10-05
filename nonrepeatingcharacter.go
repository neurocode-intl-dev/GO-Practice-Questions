package main

import "fmt"

// firstNonRepeatingCharacter returns the first character
// that appears only once in the string.
func firstNonRepeatingCharacter(s string) string {

	// Create a hash map.
	// Key   = character
	// Value = number of times the character appears
	count := make(map[rune]int)

	// First pass:
	// Count every character in the string.
	for _, char := range s {

		// Increase the count of the current character.
		count[char]++
	}

	// Second pass:
	// Go through the string again in its original order.
	for _, char := range s {

		// If the character appears only once,
		// this is the first non-repeating character.
		if count[char] == 1 {

			// Convert the rune to a string and return it.
			return string(char)
		}
	}

	// If every character repeats, return an empty string.
	return ""
}

func main() {
	s := "leetcode"
	result := firstNonRepeatingCharacter(s)
	fmt.Println(result)
}