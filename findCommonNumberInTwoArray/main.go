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
	arr := []int{2, 2}
	arr1 := []int{1, 2, 1, 4}

	set1 := make(map[int]bool)
	seen := make(map[int]bool)
	result := []int{}

	for _, v := range arr {
		set1[v] = true
	}

	for _, v := range arr1 {
		if set1[v] && !seen[v] {
			result = append(result, v)
			seen[v] = true
		}
	}

	fmt.Println(result) // Output: [2]
}
