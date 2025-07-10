package main

import (
	"fmt"
)

func canJump(arr []int) bool {
	nth := 0
	for i := range len(arr) {
		if i > nth {
			return false
		}
		if i+arr[i] > nth {
			nth = i + arr[i]
		}
	}
	return true
}

func main() {
	// Test case 1
	nums1 := []int{2, 3, 1, 1, 4}
	fmt.Println("Input:", nums1)
	fmt.Println("Can jump to end?", canJump(nums1)) // Output: true

	// Test case 2
	nums2 := []int{3, 2, 1, 0, 4}
	fmt.Println("\nInput:", nums2)
	fmt.Println("Can jump to end?", canJump(nums2)) // Output: false
}
