package main

import "fmt"

func main() {
	str := "Nikhil"

	freq := make(map[string]int)

	for _, char := range str {
		freq[string(char)] = freq[string(char)] + 1

	}

	fmt.Println("String Count:", freq)
}
