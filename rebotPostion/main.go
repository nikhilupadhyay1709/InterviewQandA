package main

import (
	"fmt"
)

func main() {
	commands := "RAARA"
	finalPosition := findFinalPosition(commands)
	fmt.Println("Final Position of the Robot:", finalPosition)
}

func findFinalPosition(commands string) int {
	position := 0 // Start at position 0
	vel := 1      // Initial velocity
	dir := 1      // 1 for forward, -1 for backward

	for _, cmd := range commands {
		switch cmd {
		case 'R':
			dir *= -1
			vel = 1
		case 'A':
			position += dir * vel
			vel *= 2
		}
	}

	return position
}
