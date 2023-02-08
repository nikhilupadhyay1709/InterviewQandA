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
	m := map[string]int{}

	for _, v := range str {
		m[string(v)] += 1
	}

	for i, v := range str {
		if m[string(v)] == 1 {
			return i
		}
	}
	return -1
}