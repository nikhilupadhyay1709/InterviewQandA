package main

import (
	"fmt"
)

func is_vowel(char rune) bool {
	if (char == 'a') || (char == 'e') || (char == 'i') ||
		(char == 'o') || (char == 'u') {
		return true
	} else {
		return false
	}
}

func count_vowels(str string) int {
	count := 0
	for _, char := range str {
		if is_vowel(char) {
			count = count + 1
		}
	}
	return count
}

func main() {
	x := count_vowels("Hello Go")
	fmt.Println("Count of Vowels :", x)

	x = count_vowels("This is golang code")
	fmt.Println("Count of Vowels :", x)

	x = count_vowels("By")
	fmt.Println("Count of Vowels :", x)
}
