package main

import "fmt"

func averageArray(arr [5]int) float64 {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return float64(sum) / float64(len(arr))
}

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Average of array elements:", averageArray(arr))
}
	