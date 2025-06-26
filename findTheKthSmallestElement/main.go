package main

import (
	"fmt"
)

// Find the Kth Smallest Element:
// Write a function that finds the k-th smallest element in a slice of integers.
func findKthSmallest(nums []int, k int) int {
	if k < 1 || k > len(nums) {
		return -1 // Invalid k
	}
	l := len(nums)

	for i := range l - 1 {
		for j := range l - i - 1 {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j] // Swap
			}
		}
	}

	return nums[k-1]
}

func main() {
	nums := []int{7, 10, 4, 3, 20, 15}
	k := 15
	result := findKthSmallest(nums, k)
	fmt.Printf("The %d-th smallest element in %v is: %d\n", k, nums, result)
}
