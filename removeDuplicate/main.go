package main

import (
	"fmt"
)

func removeDup(arr []int) []int {

	m := make(map[int]bool)

	res := []int{}

	for _, v := range arr {

		ok := m[v]

		if !ok {
			m[v] = true
			res = append(res, v)
		}
	}
	return res
}

func main() {
	arr := []int{1, 5, 3, 6, 9, 9, 4, 2, 3, 1, 5}
	fmt.Println("BeforeRemovingDuplicates :", arr)
	res := removeDup(arr)
	fmt.Println("AfterRemovingDuplicates :", res)
}
