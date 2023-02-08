package main

import (
	"fmt"
)

func appendCategory(a []string, b []string) []string {

	m := make(map[string]int)
	d := append(a, b...)
	res := make([]string, 0)
	for _, v := range d {
		m[v] = 1
	}

	for l := range m {
		res = append(res, l)
	}

	return res
}

func main() {
	a := []string{"x", "y", "z"}
	b := []string{"x", "p", "q"}
	c := appendCategory(a, b)
	fmt.Println(c)

}