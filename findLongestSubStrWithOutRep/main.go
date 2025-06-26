package main

import "fmt"

func main() {
	s := "abcabcbb"
	result := lengthOfLongestSubstring(s)
	fmt.Println(result) // Output: 3
}

// - Write a function that finds the length of the longest substring without repeating characters.
func lengthOfLongestSubstring(s string) int {
	m := make(map[int32]int)
	mlen := 0
	start := 0

	for i, v := range s {
		if i, f := m[v]; f && i >= start {
			start = i + 1
		}

		m[v] = i
		if l := i - start + 1; l > mlen {
			mlen = l
		}
	}
	return mlen
}
