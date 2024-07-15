package main

import "fmt"

func removeElement(arr []int, e int) []int {

	var res []int

	for _, v := range arr {
		if v != e {
			res = append(res, v)
		}
	}

	return res
}

func main() {

	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Array after removing :", removeElement(arr, 3))
}
