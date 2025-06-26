package main

import (
	"fmt"
	"strings"
)

// Anagram Checker:
func areAnagrams(s1, s2 string) bool {
	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)

	if len(s1) != len(s2) {
		return false
	}

	count := make(map[int32]int)

	for _, char := range s1 {
		count[char]++
	}

	for _, char := range s2 {
		if count[char] == 0 {
			return false
		}
		count[char]--
	}

	return true

}

func main() {
	s1 := "Listen"
	s2 := "Silent"
	fmt.Printf("Are '%s' and '%s' anagrams? %v\n", s1, s2, areAnagrams(s1, s2))
}
