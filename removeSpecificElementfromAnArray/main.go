package main

import "fmt"

func removeElement(arr []int, e int) []int {

	var result []int

	for _, value := range arr {
		if value != e {
			result = append(result, value)
		}
	}

	return result
}

func main() {

	arr := []int{1, 2, 3, 4, 5}

	e := 3

	fmt.Println("Array after removing", e, ":", removeElement(arr, e))
}
