package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 4, 6, 7, 2, 8, 9, 9, 45, 67}
	findPairSum(arr, 10)
}

func findPairSum(arr []int, sum int) {
	m := make(map[int]bool)

	for _, value := range arr {
		complement := sum - value
		if m[complement] {
			fmt.Println("Given Pair:", complement, value)
		}
		m[value] = true
	}
}
