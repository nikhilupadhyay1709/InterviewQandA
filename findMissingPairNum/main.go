package main

import "fmt"

func main() {
	arr := []int{1, 1, 2, 6, 3, 6, 5, 8, 9, 4, 3, 6, 5}

	freq := make(map[int]int)
	for _, num := range arr {
		freq[num]++
	}

	for num, count := range freq {
		if count%2 != 0 {
			fmt.Println("Elements without pairs:", num)
		}
	}
}
