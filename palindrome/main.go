package main

import "fmt"

func isPalindrome(x int) bool {
	sum := 0
	compare := x

	for x > 0 {
		r := x % 10
		sum = (sum * 10) + r
		x = x / 10
	}

	if compare == sum {
		return true
	}
	return false

}

func main() {
	fmt.Println("SecondFunction:", isPalindrome(121))
}
