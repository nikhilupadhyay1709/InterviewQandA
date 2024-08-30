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
	position := 0  // Start at position 0
	velocity := 1  // Initial velocity
	direction := 1 // 1 for forward, -1 for backward

	for _, command := range commands {
		if command == 'R' {
			direction *= -1
			velocity = 1
		} else if command == 'A' {
			position += direction * velocity
			velocity *= 2
		}
	}

	return position
}
