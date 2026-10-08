package main

import "fmt"

func fizzBuzz(n int) {

    for i := 1; i <= n; i++ {

        if i%3 == 0 && i%5 == 0 {
            fmt.Println("FizzBuzz")

        } else if i%3 == 0 {
            fmt.Println("Fizz")

        } else if i%5 == 0 {
            fmt.Println("Buzz")

        } else{
            fmt.Println(i)
        }
    }
	
}

func main() {

    fizzBuzz(15)
}

/*
1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
FizzBuzz

*/