package main

import "fmt"

func main() {
	str := "NITIN"
	str = "Nikhil"
	fmt.Println("Is Palindrome 🚀:", Palindrome(str))
}

func Palindrome(str string) bool {
	l := len(str)

	for i := range l / 2 {
		if str[i] != str[l-1-i] {
			return false
		}
	}

	return true
}
