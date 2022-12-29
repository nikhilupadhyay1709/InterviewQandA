package main

import "fmt"

func main() {
	fmt.Println("FindOut Max Number:", max(7, 10))
	fmt.Println("FindOut Min Number:", min(7, 10))

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
