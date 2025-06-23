// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	str := "i am Nikhil"
	res := NonRepeatingChar(str)
	fmt.Println("FirstNonRepeatingCharacter", res)
}

func NonRepeatingChar(str string) string {
	m := make(map[int32]int)

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
