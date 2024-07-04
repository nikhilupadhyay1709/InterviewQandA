package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 4, 6, 7, 2, 8, 9, 9, 45, 67}
	findPairSum(arr, 10)
}

func findPairSum(arr []int, n int) {
	m := make(map[int]bool)

	for _, v := range arr {
		com := n - v
		if m[com] {
			fmt.Println("Given Pair:", com, v)
		}
		m[v] = true
	}
}
