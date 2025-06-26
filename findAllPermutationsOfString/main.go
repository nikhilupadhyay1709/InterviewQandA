// Find All Permutations of a String:
// Write a function that generates all permutations of a given string.
package main

import (
	"fmt"
)

func permute(s string) []string {
	var res []string
	var helper func(string, string)

	helper = func(prefix, remaining string) {
		if len(remaining) == 0 {
			res = append(res, prefix)
			return
		}
		for i, char := range remaining {
			helper(prefix+string(char), remaining[:i]+remaining[i+1:])
		}
	}

	helper("", s)
	return res
}

func main() {
	s := "abc"
	permutations := permute(s)
	fmt.Println("Permutations of", s, "are:", permutations)
}
