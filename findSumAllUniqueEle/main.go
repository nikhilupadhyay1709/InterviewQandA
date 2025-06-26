package main

import "fmt"

// Sum of All Unique Elements:
// Write a function that calculates the sum of all unique elements in a slice of integers.
func sumOfUniqueElements(nums []int) int {
	m := make(map[int]int)
	sum := 0

	for _, num := range nums {
		m[num]++
	}

	for num, count := range m {
		m[num]++
		if count == 1 {
			sum += num
		}
	}
	return sum
}

func main() {
	nums := []int{1, 2, 3, 2, 4, 5, 5}
	result := sumOfUniqueElements(nums)
	fmt.Printf("The sum of unique elements in %v is: %d\n", nums, result)
}
