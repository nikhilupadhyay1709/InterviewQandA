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

/*
// This code finds the missing elements in a sorted array of integers.
// It uses a map to track the elements that exist in the array and then iterates
// through the range from the first to the last element of the array to find the missing ones.

package main

import (
	"fmt"
	"sort"
)

func missingElementsFromArray(arr []int) []int {
	sort.Ints(arr)

	result := []int{}
	start := arr[0]
	end := arr[len(arr)-1]
	exists := make(map[int]bool)

	for _, val := range arr {
		exists[val] = true
	}

	for i := start; i <= end; i++ {
		if !exists[i] {
			result = append(result, i)
		}
	}

	return result
}

func main() {
	arr := []int{1, 2, 6, 7, 10}
	res := missingElementsFromArray(arr)
	fmt.Println("Missing Numbers Are 🚀:", res)
}

*/
