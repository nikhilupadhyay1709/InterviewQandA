package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "aabcccccaaa"
	n := len(s)

	var builder strings.Builder
	count := 1

	for i := 1; i <= n; i++ {
		if i < n && s[i] == s[i-1] {
			count++
		} else {
			builder.WriteByte(s[i-1])
			builder.WriteString(strconv.Itoa(count))
			count = 1
		}
	}

	compressed := builder.String()

	if len(compressed) >= n {
		fmt.Println(s)
	} else {
		fmt.Println(compressed)
	}
}
