package main

import (
	"fmt"
	"slices"
)

func contains(arr []int, target int) bool {
	// for _, v := range arr {
	// 	if v == target {
	// 		return true
	// 	}
	// }
	// return false

	return slices.Contains(arr, target)
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	target := 3
	fmt.Println("Array contains", target, ":", contains(arr, target))
}
