// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	str := "ikhil"
	res := FirstNonRepeatingCharacter(str)
	fmt.Println("FirstNonRepeatingCharacter", res)
}

func FirstNonRepeatingCharacter(str string) int {
	charCounts := map[string]int{}

	for _, v := range str {
		charCounts[string(v)] += 1
	}

	for i, v := range str {
		if charCounts[string(v)] == 1 {
			return i
		}
	}
	return -1
}