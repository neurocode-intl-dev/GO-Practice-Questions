package main

import "fmt"

func fizzBuzz(n int) {

	// Loop from 1 up to n.
	for i := 1; i <= n; i++ {

		// Check if the number is divisible by both 3 and 5.
		// We check this first because 15 is both Fizz and Buzz.
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")

		// Check if the number is divisible by 3.
		} else if i%3 == 0 {
			fmt.Println("Fizz")

		// Check if the number is divisible by 5.
		} else if i%5 == 0 {
			fmt.Println("Buzz")

		// If it is not divisible by 3 or 5,
		// print the number itself.
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	fizzBuzz(15)
}