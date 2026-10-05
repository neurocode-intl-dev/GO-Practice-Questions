package main

import "fmt"

func isAnagram(s string, t string) bool {
	// If lengths are different, they cannot be anagrams
	if len(s) != len(t) {
		return false
	}

	// Hash map: character -> count
	// rune is a built-in type that represents a single Unicode character. Technically, it is an alias for the int32 data type.
	count := make(map[rune]int)  

	// Count characters from the first string
	for _, char := range s {
		count[char]++
	}

	// Remove characters using the second string
	for _, char := range t {
		count[char]--

		// If count becomes negative,
		// t contains a character too many times
		if count[char] < 0 {
			return false
		}
	}

	return true
}

func main() {
	s := "listen"
	t := "silent"

	result := isAnagram(s, t)

	fmt.Println(result)
}