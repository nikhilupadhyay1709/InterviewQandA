package main

import (
	"fmt"
	"strings"
)

// function to reverse string
func Reverse(str string) (res string) {
	var build strings.Builder

	defer func() {
		res = build.String()
	}()

	for _, r := range str {
		defer func(r rune) {
			build.WriteRune(r)
		}(r)
	}

	return
}

func main() {
	str := "lihkiN"
	fmt.Println(Reverse(str))

}