package main

import (
	"fmt"
)

func commonPostfix(input []string) string {

	postfix := input[0]

	for _, str := range input[1:] {
		minLength := min(len(str), len(postfix))

		for i := 1; i <= minLength; i++ {
			if postfix[len(postfix)-i] != str[len(str)-i] {
				postfix = postfix[:len(postfix)-i]
				break
			}
		}

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
