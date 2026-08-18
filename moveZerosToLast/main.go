package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 0, 2, 4, 8, 0, 9, 4, 0, 5}
	pos := 0

	for i, v := range arr {
		if v != 0 {
			arr[pos], arr[i] = arr[i], arr[pos]
			pos++
		}
	}

	fmt.Println(arr)
}
