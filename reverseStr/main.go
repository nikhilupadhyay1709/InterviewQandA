package main

import "fmt"

func main() {
	str := "Hello, Go!"

	var res string
	for _, v := range str {
		res = string(v) + res
	}

	fmt.Println("Reverse String 🚀:", res)
}
