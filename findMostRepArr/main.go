package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 3, 2, 4, 3, 2, 5, 3, 3, 5, 5, 5}

	m := make(map[int]int)
	maxCount := 0
	mostReps := []int{}

	for _, num := range arr {
		m[num]++
		count := m[num]

		if count > maxCount {
			maxCount = count
			mostReps = []int{num} // reset with new max
		} else if count == maxCount {
			mostReps = append(mostReps, num)
		}
	}

	fmt.Printf("Most Repeated Elements: %v\n", mostReps)
	fmt.Printf("Count: %d\n", maxCount)
}
