package main

import "fmt"

func isAnagram(s , t string) bool {
    if len(s) != len(t) {
        return false
    }

    count1 := map[byte]int{}
    count2 := map[byte]int{}

    for i := 0; i < len(s); i++ {
        count1[s[i]]++
    }

    for i := 0; i < len(t); i++ {
        count2[t[i]]++
    }

    //Compare both maps key value count to check if they are equal
	// i.e for same value count
    for char, freq := range count1 {
        if count2[char] != freq {
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

//true