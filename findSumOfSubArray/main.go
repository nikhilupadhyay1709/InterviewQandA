package main

import (
	"fmt"
)

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	maxSum := maxSubArray(nums)
	fmt.Println("Maximum subarray sum:", maxSum)
}

func maxSubArray(arr []int) int {
	maxSum := arr[0]
	currentSum := arr[0]

	for i := 1; i < len(arr); i++ {
		currentSum = max(arr[i], currentSum+arr[i])
		maxSum = max(maxSum, currentSum)
	}

	return maxSum
}
