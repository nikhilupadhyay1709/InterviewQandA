package main

import "fmt"

func main() {

	var n = []int{1, 1, 1, 0, 0}

	var isDone = false

	for !isDone {
		isDone = true
		var i = 0
		for i < len(n)-1 {
			if n[i] > n[i+1] {
				n[i], n[i+1] = n[i+1], n[i]

				isDone = false
			}
			i++
		}
	}

	fmt.Println("Sorted Array is 🚀:", n)
}
