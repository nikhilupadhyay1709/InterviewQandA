package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 0, 2, 4, 8, 0, 9, 4, 0, 5}
	n := len(arr)

	// Index to place the next non-zero element
	pos := n - 1

	// Move non-zero elements to the end of the array
	for i := n - 1; i >= 0; i-- {
		if arr[i] != 0 {
			arr[pos] = arr[i]
			pos--
		}
	}

	// Fill the beginning of the array with zeros
	for i := 0; i <= pos; i++ {
		arr[i] = 0
	}
	fmt.Println(arr) // Output: [0 0 0 1 2 3 4 2 4 8 9 4 5]
}
