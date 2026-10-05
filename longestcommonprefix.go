package main

import "fmt"

// longestCommonPrefix finds the longest prefix that is common to all strings.
func longestCommonPrefix(strs []string) string {

	// If the slice is empty, there is no common prefix.
	if len(strs) == 0 {
		return ""
	}

	// Use the first string as the initial prefix.
	//
	// Example:
	// strs[0] = "flower"
	//
	// prefix = "flower"
	prefix := strs[0]

	// Compare the prefix with every other string.
	for i := 1; i < len(strs); i++ {

		// Keep checking while the current string does not start with the prefix.
		for len(prefix) > 0 {

			// Check whether the current string starts with the current prefix.
			if hasPrefix(strs[i], prefix) {

				// The prefix matches, so stop reducing it for this string.
				break
			}

			// Remove the last character from the prefix.
			//
			// Example:
			// "flower" -> "flowe" -> "flow" -> ...
			prefix = prefix[:len(prefix)-1]
		}

		// If the prefix becomes empty, there is no common prefix.
		if prefix == "" {
			return ""
		}
	}

	// Return the longest prefix shared by all strings.
	return prefix
}

// hasPrefix checks whether s starts with prefix.
func hasPrefix(s string, prefix string) bool {

	// If the prefix is longer than the string, the string cannot start with the prefix.
	if len(prefix) > len(s) {
		return false
	}

	// Compare the beginning of the string with the prefix.
	return s[:len(prefix)] == prefix
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	result := longestCommonPrefix(strs)
	fmt.Println(result)
}