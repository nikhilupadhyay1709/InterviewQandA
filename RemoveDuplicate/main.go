package main

import (
	"fmt"
)

func removeDup(arr []int) []int {
	m := make(map[int]bool)

	list := []int{}

	for _, v := range arr {
		if _, ok := m[v]; !ok {
			m[v] = true
			list = append(list, v)
		}
	}
	return list
}

func main() {
	arr := []int{1, 5, 3, 6, 9, 9, 4, 2, 3, 1, 5}
	fmt.Println("BeforeRemovingDuplicates :", arr)
	res := removeDup(arr)
	fmt.Println("AfterRemovingDuplicates :", res)
}
