// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	str := "ikhil"
	res := FirstNonRepeatingCharacter(str)
	fmt.Println("FirstNonRepeatingCharacter", res)
}

func FirstNonRepeatingCharacter(str string) string {
	m := make(map[rune]int)

	for _, v := range str {
		m[v]++
	}

	for _, v := range str {
		if m[v] == 1 {
			return string(v)
		}
	}
	return ""
}
