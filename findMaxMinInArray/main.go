package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	min, max := findMaxMin(arr)

	fmt.Println("MinNum:", min, "MaxNum :", max)

}

func findMaxMin(arr []int) (int, int) {
	max := arr[0]
	min := arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		} else if arr[i] < min {
			min = arr[i]
		}

	}
	return min, max

}