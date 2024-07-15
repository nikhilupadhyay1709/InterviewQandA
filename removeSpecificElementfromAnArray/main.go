package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	e := 3

	var res []int
	for _, v := range arr {
		if v != e {
			res = append(res, v)
		}
	}

	fmt.Println("Array after removing:", res)
}
