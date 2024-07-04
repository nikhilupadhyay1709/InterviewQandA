package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "aabcccccaaa"
	n := len(s)

	var comp string
	count := 1

	for i := 1; i <= n; i++ {
		if i < n && s[i] == s[i-1] {
			count++
		} else {
			comp += string(s[i-1]) + strconv.Itoa(count)
			count = 1
		}
	}

	if len(comp) >= n {
		fmt.Println(s)
	} else {
		fmt.Println(comp)
	}
}
