package main

import "fmt"

func removeDup(str []string) []string {
	m := make(map[string]bool)
	var res []string
	for _, v := range str {
		if _, ok := m[v]; !ok {
			m[v] = true
			res = append(res, v)
		}
	}
	return res
}

func main() {
	strArr := []string{"abc", "cde", "efg", "efg", "abc", "cde"}
	res := removeDup(strArr)
	fmt.Println("AfterRemovingDuplicates :", res)
}
