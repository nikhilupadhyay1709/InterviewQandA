package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 3, 2, 1, 4, 1, 3, 3, 3, 2, 4, 4, 4, 4}

	// Map to count occurrences
	m := make(map[int]int)

	// Count frequencies
	for _, v := range arr {
		m[v]++
	}

	// Find most repeated element
	var mostRep int
	maxCount := 0
	for v, count := range m {
		if count > maxCount {
			mostRep = v
			maxCount = count
		}
	}

	fmt.Printf("Most repeated element: %d\nCount: %d\n", mostRep, maxCount)
}
