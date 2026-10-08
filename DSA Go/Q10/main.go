package main

import "fmt"

func firstNonRepeating(s string) string {

    count := map[byte]int{}

    // Count each character
    for i := 0; i < len(s); i++ {
        count[s[i]]++
    }

    // Find the first character with count 1
    for i := 0; i < len(s); i++ {

        if count[s[i]] == 1 {
            return string(s[i])
        }
    }

    return ""
}

func main() {

    text := "AAfootball"

    result := firstNonRepeating(text)

    fmt.Println(result)
}

//f