package main

import (
	"fmt"
	"strings"
)

type Position struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

func main() {
	str := []Position{
		{Start: "Home", End: "Busstop"},
		{Start: "Office", End: "Gym"},
		{Start: "Gym", End: "Home"},
		{Start: "Busstop", End: "Office"},
	}

	res := findStartToEndLocation(str, "Home")
	fmt.Println(strings.Join(res, " -> "))
}

func findStartToEndLocation(routes []Position, start string) []string {
	// Create a map from start to end
	m := make(map[string]string)
	
	for _, pos := range routes {
		m[pos.Start] = pos.End
	}

	var path []string
	current := start

	for {
		path = append(path, current)
		next, exists := m[current]
		if !exists || next == start {
			// Either route ends or comes back to start
			if exists {
				path = append(path, start)
			}
			break
		}
		current = next
	}

	return path
}
