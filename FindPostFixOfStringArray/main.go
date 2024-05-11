package main

import (
	"fmt"
)

func commonPostfix(input []string) string {
	if len(input) == 0 {
		return "" // If input array is empty, return an empty string
	}

	postfix := input[0] // Initialize postfix as the first string in the array

	// Iterate over the remaining strings in the array
	for _, str := range input[1:] {
		// Find the minimum length between current postfix and current string
		minLength := len(postfix)
		if len(str) < minLength {
			minLength = len(str)
		}

		// Compare characters from the end to find the common postfix
		for i := 1; i <= minLength; i++ {
			if postfix[len(postfix)-i] != str[len(str)-i] {
				postfix = postfix[:len(postfix)-i] // Trim postfix if characters don't match
				break
			}
		}

		// If postfix becomes empty, break the loop as there's no common postfix
		if postfix == "" {
			break
		}
	}

	return postfix
}

func main() {
	input := []string{"alee", "pee", "apriee"}
	result := commonPostfix(input)
	fmt.Println("Common postfix:", result)
}
