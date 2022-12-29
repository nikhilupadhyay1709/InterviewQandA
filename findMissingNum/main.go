package main

import (
	"fmt"
	"strconv"
)

// FindMissing function will find the missing number in an inpur array
func FindMissing(input []int) int {

	actualSum := 0

	for _, inp := range input {
		actualSum += inp
	}

	n := len(input)
	sumOfN := float64(n*(n+1)) / float64(2)

	return int(sumOfN - float64(actualSum))
}

func ArrayToString(nums []int) string {
	if len(nums) == 0 {
		return "[]"
	}
	res := "["
	for _, num := range nums {
		res += strconv.Itoa(num) + ", "
	}
	res += "]"
	return res[:len(res)-3] + "]"
}

func main() {
	array := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 14, 16, 17, 19},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 14, 16, 17, 19},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 14, 16, 17, 19},
	}
	for i, arr := range array {
		fmt.Printf("%d.     Original: %s\n", i+1, ArrayToString(arr))
		missingNumber := FindMissing(arr)
		fmt.Printf("       The missing number is: %d\n", missingNumber)
	}
}
