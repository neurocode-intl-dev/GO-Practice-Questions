package main

import "fmt"



func isValid(s string) bool {

	// Stack is used to store the opening brackets of type byte
	stack := []byte{}

	for i := 0; i < len(s); i++ {

		char := s[i]

		if char == '(' || char == '[' || char == '{' {

			stack = append(stack, char)

		} else {

			if len(stack) == 0 {
				return false
			}

			last := stack[len(stack)-1]

			if char == ')' && last != '(' {
				return false
			}

			if char == ']' && last != '[' {
				return false
			}

			if char == '}' && last != '{' {
				return false
			}

			// Remove last item from stack
			stack = stack[:len(stack)-1]
			// fmt.Println(stack)
		}
	}

	return len(stack) == 0
}


func main() {

	fmt.Println(isValid("()"))
	fmt.Println(isValid("([{}])"))
	fmt.Println(isValid("([)]"))
}

// []
// true
// [40 91]
// [40]
// []

// true
// true
// false