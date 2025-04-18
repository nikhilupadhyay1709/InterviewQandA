package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	max, min := findMaxAndMin(arr)

	fmt.Printf("maxNum: %d, minNum: %d\n", max, min)
}

func findMaxAndMin(arr []int) (int, int) {
	max := arr[0]
	min := arr[0]

	for i := range arr {
		if arr[i] > max {
			max = arr[i]
		} else {
			min = arr[i]
		}
	}

	return max, min
}
