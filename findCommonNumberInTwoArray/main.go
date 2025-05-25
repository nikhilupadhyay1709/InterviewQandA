package main

import "fmt"

func main() {
	arr := []int{2, 2}
	arr1 := []int{1, 2, 1, 4}

	set := make(map[int]bool)
	seen := make(map[int]bool)
	res := []int{}

	for _, v := range arr {
		set[v] = true
	}

	for _, v := range arr1 {
		if set[v] && !seen[v] {
			res = append(res, v)
			seen[v] = true
		}
	}

	fmt.Println(res) // Output: [2]
}
