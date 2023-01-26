package main

import "fmt"

func main() {
	str := "()"
	res := isValid(str)
	fmt.Println("isValidParenthesis:", res)
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	st := []rune{}
	open := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for _, r := range s {
		if closer, ok := open[r]; ok {
			st = append(st, closer)
			continue
		}

		l := len(st) - 1
		if l < 0 || r != st[l] {
			return false
		}

		st = st[:l]
	}

	return len(st) == 0
}