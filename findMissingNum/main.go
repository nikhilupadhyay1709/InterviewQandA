package main

import (
	"fmt"
	"sort"
)

func missingElementsFromArray(arr []int, N int) []int {

	// sort the array
	sort.Ints(arr)

	result := []int{}
	var diff = arr[0] - 0

	for i := range N {
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