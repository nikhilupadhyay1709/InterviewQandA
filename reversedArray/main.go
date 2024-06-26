package main

import "fmt"

func reverseArray(arr []int) {
	length := len(arr)

	for i := 0; i < length/2; i++ {
		arr[i], arr[length-1-i] = arr[length-1-i], arr[i]
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Original array:", arr)
	reverseArray(arr)

	fmt.Println("Reversed array:", arr)
}
