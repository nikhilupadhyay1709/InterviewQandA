package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	// Sort intervals based on the starting time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := merged[len(merged)-1]
		current := intervals[i]

		// Check for overlap
		if current[0] <= last[1] {
			last[1] = max(last[1], current[1])
		} else {
			merged = append(merged, current)
		}
	}

	return merged
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println("Merged intervals:", merge(intervals))

	intervals2 := [][]int{{1, 4}, {4, 5}}
	fmt.Println("Merged intervals:", merge(intervals2))
}
