package main

import (
	"fmt"
	"strconv"
)

func compressString(s string) string {
	n := len(s)

	var compressed string
	count := 1

	for i := 1; i <= n; i++ {
		if i < n && s[i] == s[i-1] {
			count++
		} else {
			compressed += string(s[i-1]) + strconv.Itoa(count)
			count = 1
		}
	}

	if len(compressed) >= n {
		return s
	}
	return compressed
}

func main() {
	testString := "aabcccccaaa"
	result := compressString(testString)
	fmt.Println(result)
}
