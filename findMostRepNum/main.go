package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 2, 3, 2, 1, 12}
	m := make(map[int]int)
	maxCount := 0
	
	// Count frequencies and track max count
	for _, v := range arr {
		m[v]++
		if m[v] > maxCount {
			maxCount = m[v]
		}
	}

	// Collect numbers with the highest frequency
	mostFrequent := []int{}
	for num, count := range m {
		if count == maxCount {
			mostFrequent = append(mostFrequent, num)
		}
	}

	fmt.Printf("Most repeated number(s) 🚀: %v (Count: %d)\n", mostFrequent, maxCount)

}
