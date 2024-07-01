package main

import "fmt"

func main() {
	var year int
	
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		fmt.Println("The year is a leap year!")
	} else {
		fmt.Println("The year isn't a leap year!")
	}
}
