package main

import "fmt"

func removeElement(arr []int, element int) []int {

	var result []int

	for _, value := range arr {
		if value != element {
			result = append(result, value)
		}
	}

	return result
}

func main() {

	arr := []int{1, 2, 3, 4, 5}

	element := 3

	fmt.Println("Array after removing", element, ":", removeElement(arr, element))
}
