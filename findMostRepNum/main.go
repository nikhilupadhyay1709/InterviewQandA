package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 10, 11, 11, 12}
	freq := make(map[int]int)
	maxnum := 0

	var mostRepeated, maxCount = 0, 0
	for _, val := range arr {
		freq[val] = freq[val] + 1
	}

	for num, count := range freq {
		if count > maxnum {
			maxnum = count
		}
		if maxnum == count {
			maxCount = maxnum
			mostRepeated = num
		}
	}
	fmt.Println("Most repeated Number is 🚀:", mostRepeated, "Count of 🚀:", maxCount)
}
