package main

import "fmt"

func main() {
	arr := []int{1, 12, 13, 14, 4}
	arr1 := []int{9, 8, 1, 4, 13}

	m := make(map[int]int)
	combined := append(arr, arr1...)

	for _, v := range combined {
		m[v]++
		if m[v] == 2 {
			fmt.Println("Common Number:", v)
		}
	}
}
