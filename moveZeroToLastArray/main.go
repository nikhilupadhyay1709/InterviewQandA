package main

import "fmt"

func main() {
	arr := []int{0, 1, 0, 3, 12}

	// Move all zeros to the end of the array
	pos := 0
	for _, v := range arr {
		if v != 0 {
			arr[pos] = v
			pos++
		}
	}
	for pos < len(arr) {
		arr[pos] = 0
		pos++
	}

	fmt.Println(arr) // Output: [1 3 12 0 0]
}
