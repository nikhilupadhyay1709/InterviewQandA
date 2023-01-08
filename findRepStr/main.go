package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Nikhil"

	fmt.Println("Full String:", str)

	count := strings.Count(str, "i")
	fmt.Println("The count of i is 🚀:", count)
}