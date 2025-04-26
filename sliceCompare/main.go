package main

import (
	"bytes"
	"fmt"
)

func main() {
	a := []byte{1, 2, 3}
	b := []byte{1, 2, 3}
	c := []byte{1, 2, 4}

	// Compare a and b
	if bytes.Equal(a, b) {
		fmt.Println("a and b are equal")
	} else {
		fmt.Println("a and b are not equal")
	}

	// Compare a and c
	if bytes.Equal(a, c) {
		fmt.Println("a and c are equal")
	} else {
		fmt.Println("a and c are not equal")
	}
}
