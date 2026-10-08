package main

import "fmt"

func reverseString(word string) string {

    result := ""


    for i := len(word) - 1; i >= 0; i-- {
        result = result + string(word[i])
    }

    return result
}

func main() {

    text := "Computer"
    result := reverseString(text)

    fmt.Println(result)
}

//retupmoC