package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{1, 2, 3, 10}
	secondMinNum := findSecondMinNum(arr)
	if secondMinNum == math.MaxInt64 {
		fmt.Println("No second minimum value found")
	} else {
		fmt.Println("secondMinNum:", secondMinNum)
	}
}

func findSecondMinNum(arr []int) int {
	if len(arr) < 2 {
		return math.MaxInt64
	}

	min, secondMin := math.MaxInt64, math.MaxInt64

	for _, num := range arr {
		if num < min {
			secondMin, min = min, num
		} else if num < secondMin && num != min {
			secondMin = num
		}
	}

	return secondMin
}
