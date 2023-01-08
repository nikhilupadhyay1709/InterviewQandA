package main

import "fmt"

func main() {
	x := []int{2, 3, 2, 4, 5, 67, 8}
	swap(x)
	fmt.Println("Swap Value 🚀:", x)
}
func swap(s []int) {
	for a, b := 0, len(s)-1; a < b; a, b = a+1, b-1 {
		s[a], s[b] = s[b], s[a]
	}
}
