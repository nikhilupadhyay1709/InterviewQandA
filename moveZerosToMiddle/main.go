package main

import (
	"fmt"
	"slices"
)

func moveZerosToMiddle(nums []int) []int {
	// Step 1: Collect non-zero elements and count zeros
	var nonZeros []int
	zeroCount := 0

	for _, num := range nums {
		if num == 0 {
			zeroCount++
		} else {
			nonZeros = append(nonZeros, num)
		}
	}

	// Step 2: Calculate the middle position
	mid := len(nonZeros) / 2
	left := nonZeros[:mid]
	right := nonZeros[mid:]

	// Step 3: Build result with zeros in the middle
	result := slices.Clone(left)
	for range zeroCount {
		result = append(result, 0)
	}
	result = append(result, right...)

	return result
}

func main() {
	slice := []int{1, 0, 2, 0, 3, 0, 4, 5}
	result := moveZerosToMiddle(slice)
	fmt.Println(result) // Output: [1 2 0 0 0 3 4 5]
}
