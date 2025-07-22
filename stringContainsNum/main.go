package main

import (
	"fmt"
)

func main() {
	word := "this is 23 nikhil 17 updhaay89 here for interview314"
	num := getNumsInString(word)
	fmt.Println("array before the sum :", num)
	res := sum(num)
	fmt.Println("res after the sum", res)
}

func sum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}

	return sum
}

func getNumsInString(s string) []int {
	var arr []int
	for _, v := range s {
		if v >= '0' && v <= '9' {
			arr = append(arr, int(v-'0'))
		}
	}
	return arr
}
