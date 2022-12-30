package main

import "fmt"

func duplicateInArray(arr []int) int {
	visited := make(map[int]bool, 0)
	for i := 0; i < len(arr); i++ {
		if visited[arr[i]] == true {
			return arr[i]
		} else {
			visited[arr[i]] = true
		}
	}
	return 0
}

func main() {
	fmt.Println(duplicateInArray([]int{1, 4, 7, 2, 2}))
	fmt.Println(duplicateInArray([]int{1, 4, 7, 2, 3}))
}
