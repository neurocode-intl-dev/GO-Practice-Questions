package main

import "fmt"

// countVowels counts the number of vowels in a string.
func countVowels(s string) int {

	// Create a variable to store the vowel count.
	count := 0

	// Loop through each character in the string.
	for _, char := range s {

		// Check whether the character is a vowel.
		if char == 'a' ||
			char == 'e' ||
			char == 'i' ||
			char == 'o' ||
			char == 'u' {

			// If it is a vowel, increase the count.
			count++
		}
	}

	// Return the total number of vowels.
	return count
}

func main() {
	s := "hello world"
	result := countVowels(s)
	fmt.Println(result)
}