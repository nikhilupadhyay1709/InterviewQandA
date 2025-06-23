package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{1, 2, 3, 10}

	min, secondMin := math.MaxInt64, math.MaxInt64

	for _, v := range arr {
		if v < min {
			secondMin, min = min, v
		} else if v < secondMin && v != min {
			secondMin = v
		}
	}

	fmt.Printf("The second minimum number in the array is: %d\n", secondMin)
}
