package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 1, 2, 3}
	seen := make(map[int]bool)

	// i is the write index
	i := 0
	for _, v := range arr {
		if !seen[v] {
			seen[v] = true
			arr[i] = v
			i++
		}
	}

	// Slice the array to the new length
	arr = arr[:i]

	fmt.Println(arr) // Output: [1 2 3 4 5 6]
}
