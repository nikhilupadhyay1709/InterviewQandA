package main

import (
	"fmt"
	"regexp"
)

func main() {
	word := "test25"
	numeric := regexp.MustCompile(`\d`).MatchString(word)
	fmt.Println(numeric)
}

