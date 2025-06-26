package main

func main() {
	s := "hello world, hello universe"
	sub := "hello"
	count := countOccurrences(s, sub)
	println(count) // Output: 2
}

// Count Occurrences of a Substring:
// -   Write a function that counts how many times a substring appears in a string.
func countOccurrences(s, sub string) int {
	count := 0
	subLen := len(sub)
	for i := 0; i <= len(s)-subLen; i++ {
		if s[i:i+subLen] == sub {
			count++
		}
	}
	return count
}
