package main

import "fmt"

// Function to find the sum of all elements in an array
func main() {
	arr := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, value := range arr {
		sum += value
	}

	fmt.Println("Sum of array elements:", sum)
}
