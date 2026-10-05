package main

import "fmt"

// isValid checks whether the brackets in the string
// are correctly opened and closed.
func isValid(s string) bool {

	// Create a slice to use as a Stack.
	// We will store opening brackets:
	// (, [, {
	stack := []rune{}

	// Loop through every character in the string.
	for _, char := range s {

		// Check if the character is an opening bracket.
		if char == '(' || char == '[' || char == '{' {

			// Push the opening bracket onto the stack.
			stack = append(stack, char)

		} else {

			// Before accessing the top of the stack, make sure the stack is not empty.
			//
			// Example:
			// Input = ")"
			// There is no opening bracket to match it.
			if len(stack) == 0 {
				return false
			}

			// Get the last element from the stack.
			// This is the most recently added opening bracket.
			top := stack[len(stack)-1]

			// Remove the top element from the stack.
			// This is the POP operation.
			stack = stack[:len(stack)-1]

			// Check whether the opening bracket matches the current closing bracket.
			if char == ')' && top != '(' {
				return false
			}

			if char == ']' && top != '[' {
				return false
			}

			if char == '}' && top != '{' {
				return false
			}
		}
	}

	// If the stack is empty, every opening bracket had a matching closing bracket.
	// If the stack is not empty, some opening brackets were never closed.
	return len(stack) == 0
}

func main() {
	s := "([{}])"
	result := isValid(s)
	fmt.Println(result)
}