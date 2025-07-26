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
	sum := arr[0]
	maxSum := arr[0]

	for i := range len(arr) {
		sum = max(arr[i], sum+arr[i])
		maxSum = max(maxSum, sum)
	}

	return maxSum
}
