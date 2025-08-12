package main

import "fmt"

func main() {
	str := "Golang is fast and Golang is powerful"

	freq := make(map[string]int)

	for _, v := range str {
		if v == ' ' {
			continue
		}
		freq[string(v)] = freq[string(v)] + 1
	}

	fmt.Println("Word frequencies:", freq)
}
