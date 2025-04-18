package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	l := len(arr)

	for i := range l/2 {
		arr[i], arr[l-1-i] = arr[l-1-i], arr[i]
	}

	fmt.Println(arr)
}
