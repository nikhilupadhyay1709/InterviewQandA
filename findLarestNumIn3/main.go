package main

import (
	"fmt"
)

func main() {
	a := 12
	b := 25
	c := 9

	l := a

	if b > l {
		l = b
	}
	if c > l {
		l = c
	}

	fmt.Print(l)
}
