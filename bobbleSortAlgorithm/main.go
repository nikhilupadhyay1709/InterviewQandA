package main

import "fmt"

func bobbleSort(arr []int) []int {
	l := len(arr)

	for i := range l - 1 {
		for j := range l - i - 1 {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Sorted array is:", bobbleSort(arr))
}
