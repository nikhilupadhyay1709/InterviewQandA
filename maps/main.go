package main

import (
	"fmt"
)

func main() {
	m := make(map[string]string)

	m["Gender"] = "Male"
	m["Age"] = "23"
	m["Name"] = "Nikhil"
	m["lastName"] = "Upadhyay"

	fmt.Println("name :", m["Name"], m["lastName"], "Age is:", m["Age"], "Gender is:", m["Gender"])

	delete(m, "Age")

	if len(m) == 0 {
		fmt.Println("Map is empty")
	} else {
		fmt.Println("Map has value")
	}
	// fmt.Println("Map After Deleting Age :", m)

}
