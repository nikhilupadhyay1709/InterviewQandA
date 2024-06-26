package main

import "fmt"

func sumArray(arr []int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Sum of array elements:", sumArray(arr))
}
