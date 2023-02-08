package main

import "fmt"

func findSumPair(arr []int, sum int) {
	mapping := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		if mapping[sum-arr[i]] == 0 {
			mapping[arr[i]] = i
		} else {
			fmt.Printf("Pair for given sum is (%d, %d).\n", arr[mapping[sum-arr[i]]], arr[i])
			return
		}
	}
}

func main() {
	findSumPair([]int{4, 3, 6, 7, 8, 1, 9}, 15)
	findSumPair([]int{4, 3, 6, 7, 8, 1, 9}, 100)
}
