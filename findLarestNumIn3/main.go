package main

import (
	"fmt"
)

func main() {
	a := 12
	b := 25
	c := 9

	largest := a

	if b > largest {
		largest = b
	}
	if c > largest {
		largest = c
	}

	fmt.Print(largest)
}
