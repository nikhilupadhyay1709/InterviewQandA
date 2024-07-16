package main

import "fmt"

func main() {
	fmt.Println("Max Number 🚀:", max(7, 10))
	fmt.Println("Min Number 🚀:", min(7, 10))

}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
