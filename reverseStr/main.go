package main

import "fmt"

// function to reverse string
func reverseString(str string) (res string) {
	// iterate over str and prepend to result
	for _, v := range str {
		res = string(v) + res
	}

	return
}

func main() {
	str := "Nikhil"

	fmt.Println("Original String 🚀:", str)
	// invoke reverseString
	fmt.Println("Reverse String 🚀:", reverseString(str))

}
