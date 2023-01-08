package main

import (
	"fmt"
)

func missingElementsFromArray(arr []int, N int) []int {
	result := []int{}
	var diff = arr[0] - 0

	for i := 0; i < N; i++ {
		if arr[i]-i != diff {
			for diff < arr[i]-i {
				result = append(result, i+diff)
				diff++
			}
		}
	}

	return result
}

func main() {

	arr := []int{1, 2, 6, 7, 10}
	var N = len(arr)

	res := missingElementsFromArray(arr, N)
	fmt.Println("Missing Numbers Are 🚀:", res)

}