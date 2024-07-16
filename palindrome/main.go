package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(n int) bool {
	str := strconv.Itoa(n)

	l := len(str)

	// Compare the string with its reverse
	for i := 0; i < l/2; i++ {
		if str[i] != str[l-1-i] {
			return false
		}
	}
	return true
}

func main() {
	number := 121

	if isPalindrome(number) {
		fmt.Printf("%d is a palindrome\n", number)
	} else {
		fmt.Printf("%d is not a palindrome\n", number)
	}
}