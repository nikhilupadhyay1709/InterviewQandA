package main

import "fmt"

func unique(s []string) []string {
	m := make(map[string]bool)
	var res []string
	for _, v := range s {
		if _, ok := m[v]; !ok {
			m[v] = true
			res = append(res, v)
		}
	}
	return res
}

func main() {
	fmt.Println(unique([]string{"abc", "cde", "efg", "efg", "abc", "cde"}))
}
