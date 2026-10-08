package main

import "fmt"

func longestCommonPrefix(words []string) string {

	if len(words) == 0 {
		return ""
	}

	prefix := ""



	for i := 0; i < len(words[0]); i++ {

		//char is the current character in ascii
		char := words[0][i]

		fmt.Println(string(char))

		//loop through the rest of the words
		for j := 1; j < len(words); j++ {

			if i >= len(words[j]) || words[j][i] != char {
				return prefix
			}
		}

		prefix = prefix + string(char)
	}

	return prefix
}

func main() {

	words := []string{"flower", "flow", "flight"}

	result := longestCommonPrefix(words)

	fmt.Println(result)
}

// 102 - f
// 108 - l
// 111 - o