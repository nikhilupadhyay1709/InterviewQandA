package main

import (
	"fmt"
)

func main() {
	n := 29
	checkPrime(n)
}

func checkPrime(n int) {
	if n <= 1 {
		fmt.Println(n, "is not a prime number")
		return
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			fmt.Println(n, "is not a prime number")
			return
		}
	}
	fmt.Println(n, "is a prime number")
}
