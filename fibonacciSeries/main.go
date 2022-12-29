//GOLANG Program to Generate Fibonacci Sequence Up to a Certain Number

package main

import "fmt"

func main() {
	var n int = 12
	t1 := 0
	t2 := 1
	nextTerm := 0

	for i := 1; i <= n; i++ {
		if i == 1 {
			fmt.Print(" ", t1)
			continue
		}
		if i == 2 {
			fmt.Print(" ", t2)
			continue
		}
		nextTerm = t1 + t2
		t1 = t2
		t2 = nextTerm
		fmt.Print(" ", nextTerm)
	}
}
