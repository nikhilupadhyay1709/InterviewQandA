package main

import (
	"fmt"
	"strconv"
)

func main() {

	str1 := "123"

	// using ParseInt method
	int1, err := strconv.ParseInt(str1, 6, 12)

	fmt.Println(int1)

	// If the input string contains the integer
	// greater than base, get the zero value in return
	str2 := "123"

	int2, err := strconv.ParseInt(str2, 2, 32)
	fmt.Println(int2, err)

}