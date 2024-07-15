package main

import (
	"fmt"
)

func findLongestConsecutiveSequence(arr []int) ([]int, int) {
	if len(arr) == 0 {
		return []int{}, 0
	}

	numMap := make(map[int]bool)
	for _, num := range arr {
		numMap[num] = true
	}

	longestStreak := 0
	startOfLongest := 0

	for _, num := range arr {
		if !numMap[num-1] {
			currentNum := num
			currentStreak := 1

			for numMap[currentNum+1] {
				currentNum++
				currentStreak++
			}

			if currentStreak > longestStreak {
				longestStreak = currentStreak
				startOfLongest = num
			}
		}
	}

	longestSequence := make([]int, longestStreak)
	for i := 0; i < longestStreak; i++ {
		longestSequence[i] = startOfLongest + i
	}

	return longestSequence, startOfLongest + longestStreak - 1
}

func main() {
	arr := []int{1, 94, 93, 1000, 5, 92, 78}
	longestSequence, highestElement := findLongestConsecutiveSequence(arr)
	fmt.Println("Longest consecutive sequence:", longestSequence)
	fmt.Println("Highest element in the sequence:", highestElement)
}
