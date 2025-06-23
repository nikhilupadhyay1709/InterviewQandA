package main

import (
	"fmt"
)

func findLongestConsecutiveSequence(arr []int) ([]int, int) {
	var longestSequence []int
	numMap := make(map[int]bool)
	
	longestStreak := 0

	for _, num := range arr {
		numMap[num] = true
	}

	for _, num := range arr {
		if !numMap[num-1] {
			currentNum := num
			currentStreak := 1
			currentSequence := []int{currentNum}

			for numMap[currentNum+1] {
				currentNum++
				currentStreak++
				currentSequence = append(currentSequence, currentNum)
			}

			if currentStreak > longestStreak {
				longestStreak = currentStreak
				longestSequence = currentSequence
			}
		}
	}

	return longestSequence, longestSequence[len(longestSequence)-1]
}

func main() {
	arr := []int{1, 94, 93, 1000, 5, 92, 78}
	longestSequence, highestElement := findLongestConsecutiveSequence(arr)
	fmt.Println("Longest consecutive sequence:", longestSequence)
	fmt.Println("Highest element in the sequence:", highestElement)
}
