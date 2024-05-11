package main

import "fmt"

func main() {
	input := []int{1, 12, 13, 14, 4}
	input1 := []int{9, 8, 1, 4, 13}

	// Create a map to store occurrences of numbers in the first array
	numMap := make(map[int]bool)
	for _, num := range input {
		numMap[num] = true
	}

	// Iterate through the second array and check for common numbers
	fmt.Println("Common numbers:")
	for _, num := range input1 {
		if numMap[num] {
			fmt.Println(num)
		}
	}
}
