package main

import "fmt"

func main() {
	str := "NITIN"
	str = "Nikhil"
	fmt.Println("Is Palindrome 🚀:", IsPalindrome(str))
}

func IsPalindrome(str string) bool {
	res := []byte{}

	for i := len(str) - 1; i >= 0; i-- {
		res = append(res, str[i])
	}

	return str == string(res)
}
