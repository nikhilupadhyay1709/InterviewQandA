package main

import "fmt"

func main() {
	arr := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(findMaxElement(arr))
}

func findMaxElement(arr []int) int {
	max_num := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max_num {
			max_num = arr[i]
		}
	}
	return max_num
}
