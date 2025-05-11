package main

import "fmt"

func main() {
	arr := []int{1, 12, 13, 14, 4}
	arr1 := []int{9, 8, 1, 4, 13}

	m := make(map[int]int)
	com := append(arr, arr1...)

	for _, v := range com {
		m[v]++
		if m[v] == 2 {
			fmt.Println("Common Number:", v)
		}
	}
}
func anotherWay() {
	arr := []int{1, 12, 13, 14, 4}
	arr1 := []int{9, 8, 1, 4, 13}

	seen := make(map[int]bool)
	common := make(map[int]bool)

	for _, v := range arr {
		seen[v] = true
	}

	for _, v := range arr1 {
		if seen[v] && !common[v] {
			fmt.Println("Common Number:", v)
			common[v] = true
		}
	}
}
