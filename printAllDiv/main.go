package main

import (
	"fmt"
	"math"
)

// Function to print all divisors of a given number
func printDivisors(number int) {
	fmt.Printf("Divisors of %d are:\n", number)

	for i := 1; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			// If i is a divisor, print it
			fmt.Println(i)

			// If i is not the square root of number, print the complementary divisor
			if i != number/i {
				fmt.Println(number / i)
			}
		}
	}
}

func main() {
	// Define the number
	number := 36

	// Print all divisors of the number
	printDivisors(number)
}
