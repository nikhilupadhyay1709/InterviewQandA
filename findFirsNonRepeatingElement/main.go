package main

import (
	"fmt"
)

func firstNonRepeating(arr []int) int {
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}

	for _, v := range arr {
		if m[v] == 1 {
			return v
		}
	}

	return -1
}

func main() {
	arr := []int{4, 5, 1, 2, 0, 4}
	result := firstNonRepeating(arr)
	fmt.Printf("First Non-Repeating Element: %d\n", result)
}
