package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 0, 2, 4, 8, 0, 9, 4, 0, 5}
	n := len(arr)
	pos := n - 1

	for i := n - 1; i >= 0; i-- {
		if arr[i] != 0 {
			arr[pos], arr[i] = arr[i], arr[pos]
			pos--
		}
	}

	fmt.Println(arr)
}
