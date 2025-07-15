package main

import (
	"fmt"
)

func firstNonRepeating(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	for _, num := range nums {
		if m[num] == 1 {
			return num
		}
	}

	return -1
}

func main() {
	arr := []int{4, 5, 1, 2, 0, 4}
	result := firstNonRepeating(arr)
	fmt.Printf("First Non-Repeating Element: %d\n", result)
}
