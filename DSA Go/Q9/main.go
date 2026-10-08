package main

import "fmt"

func countVowels(s string) int {
    
	count := 0


    for _ , char := range s {
        switch char {
        case 'a','e','i','o','u','A','E','I','O','U':
            count++
        }
    }
    return count
}

func main() {

    text := "democracy"

    result := countVowels(text)

    fmt.Println(result)
}

//3