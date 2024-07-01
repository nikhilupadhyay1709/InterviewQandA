package main

import (
	"fmt"
	"sort"
)

// Nested Array Flattening and Sorting

// Function to flatten a nested array
func flatten(input interface{}) []float64 {
	var flat []float64

	switch v := input.(type) {
	case []interface{}:
		for _, item := range v {
			flat = append(flat, flatten(item)...)
		}
	case []float64:
		flat = append(flat, v...)
	case float64:
		flat = append(flat, v)
	default:
		fmt.Printf("Unsupported type: %T\n", v)
	}

	return flat
}

// Function to sort the flattened array
func sortNestedArray(nestedArray []interface{}) []float64 {
	// Flatten the nested array
	flatArray := flatten(nestedArray)

	// Sort the flattened array
	sort.Float64s(flatArray)

	return flatArray
}

func main() {
	// Example nested array
	nestedArray := []interface{}{
		4.0, 1.0, []interface{}{1.0, 3.0}, 2.0, []interface{}{2.0, 7.0, []interface{}{-6.0, 0.0}},
	}

	// Print the nested array before sorting
	fmt.Printf("Before Sorting: %v\n", nestedArray)

	// Sort the nested array
	sortedArray := sortNestedArray(nestedArray)

	// Print the sorted flat array
	fmt.Printf("After Sorting: %v\n", sortedArray)
}
