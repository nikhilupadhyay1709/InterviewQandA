package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 10, 11, 11, 12}
	m := make(map[int]int)
	maxNum := 0

	var mostRepeated, maxCount = 0, 0
	for _, val := range arr {
		m[val] = m[val] + 1
	}

	for num, count := range m {
		if count > maxNum {
			maxNum = count
		}
		if maxNum == count {
			maxCount = maxNum
			mostRepeated = num
		}
	}
	fmt.Println("Most repeated Number is 🚀:", mostRepeated, "Count of 🚀:", maxCount)
}
