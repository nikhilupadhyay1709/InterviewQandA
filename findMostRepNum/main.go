package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 2, 3, 2, 1, 12}
	m := make(map[int]int)

	// Count occurrences of each number
	for _, v := range arr {
		m[v]++
	}

	// Find the most repeated number
	mostRepeated, maxCount := 0, 0
	for num, count := range m {
		if count > maxCount {
			maxCount = count
			mostRepeated = num
		}
	}

	fmt.Println("Most repeated Number is 🚀:", mostRepeated, "Count of 🚀:", maxCount)
}
