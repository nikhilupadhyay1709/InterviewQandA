package main

import "fmt"

func main() {
	input := []string{"ale", "ape", "april"}

	if len(input) == 0 {
		fmt.Println("")
		return
	}

	prefix := input[0]

	for i := 1; i < len(input); i++ {
		minLength := min(len(prefix), len(input[i]))

		prefix = prefix[:minLength]

		for j := 0; j < minLength; j++ {
			if prefix[j] != input[i][j] {
				prefix = prefix[:j]
				break
			}
		}

		if prefix == "" {
			break
		}
	}

	if prefix == "" {
		fmt.Println("No common prefix found")
	} else {
		fmt.Println(prefix)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
