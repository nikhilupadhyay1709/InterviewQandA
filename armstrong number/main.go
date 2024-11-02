package main

import (
	"fmt"
	"math"
)

// Function to check if a number is an Armstrong number
func isArmstrong(number int) bool {
	sum := 0
	temp := number
	numDigits := int(math.Log10(float64(number)) + 1)

	for temp > 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(numDigits)))
		temp /= 10
	}

	return sum == number
}

// Function to print Armstrong numbers in a given range
func findArmstrongInRange(start, end int) {
	fmt.Printf("Armstrong numbers between %d and %d are:\n", start, end)
	for i := start; i <= end; i++ {
		if isArmstrong(i) {
			fmt.Println(i)
		}
	}
}

func main() {
	// Define the range
	start := 100
	end := 999

	// Find and print Armstrong numbers in the range
	findArmstrongInRange(start, end)
}
