package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	l := len(arr)

	m := l / 2

	if l%2 != 0 {
		fmt.Printf("%d\n", arr[m])
	} else {
		fmt.Printf("%d and %d\n", arr[m-1], arr[m])
	}
}
