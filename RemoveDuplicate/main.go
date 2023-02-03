package main

import (
	"fmt"
)

func unique(intSlice []int) []int {
	m := make(map[int]bool)
	
	list := []int{}
	
	for _, v := range intSlice {
		if _, ok := m[v]; !ok {
			m[v] = true
			list = append(list, v)
		}
	}
	return list
}

func main() {
	intSlice := []int{1, 5, 3, 6, 9, 9, 4, 2, 3, 1, 5}
	fmt.Println(intSlice)
	uniqueSlice := unique(intSlice)
	fmt.Println(uniqueSlice)
}
