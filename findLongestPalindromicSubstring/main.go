package main

import (
	"fmt"
)

// Longest Palindromic Substring:
// Write a function that finds the longest palindromic substring in a given string.

func longestPalindromicSubstring(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}

	start, end := 0, 0

	for i := 0; i < n; i++ {
		len1 := expandAroundCenter(s, i, i)   // Odd length palindrome
		len2 := expandAroundCenter(s, i, i+1) // Even length palindrome
		maxLen := max(len1, len2)

		if maxLen > end-start {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}

	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func main() {
	s := "babad"
	result := longestPalindromicSubstring(s)
	fmt.Printf("The longest palindromic substring in '%s' is: '%s'\n", s, result)
}
