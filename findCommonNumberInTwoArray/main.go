package main

import "fmt"

func main() {
	arr := []int{1, 12, 13, 14, 4}
	arr1 := []int{9, 8, 1, 4, 13}

	m := make(map[int]int)
	com := append(arr, arr1...)

	for _, v := range com {
		m[v]++
		if m[v] == 2 {
			fmt.Println("Common Number:", v)
		}
	}

	/*
		// Alternative method using map to find common elements
		// This method is more efficient for larger arrays
		// as it avoids nested loops and reduces time complexity.
		// It uses a single loop to create a map of elements from the first array,
		// and then checks for common elements in the second array.
		// This is a more efficient way to find common elements in two arrays.
		// It has a time complexity of O(n + m) where n and m are the lengths of the two arrays.
		// The space complexity is O(n) for the map used to store elements of the first array.
		// This is a more efficient way to find common elements in two arrays.
		// It has a time complexity of O(n + m) where n and m are the lengths of the two arrays.
		// The space complexity is O(n) for the map used to store elements of the first array.
		// This is a more efficient way to find common elements in two arrays.
			arr := []int{1, 12, 13, 14, 4}
			arr1 := []int{9, 8, 1, 4, 13}

			seen := make(map[int]bool)
			common := make(map[int]bool)

			for _, v := range arr {
				seen[v] = true
			}

			for _, v := range arr1 {
				if seen[v] && !common[v] {
					fmt.Println("Common Number:", v)
					common[v] = true
				}
			}
	*/
}
