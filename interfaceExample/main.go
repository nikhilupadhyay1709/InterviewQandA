package main

import (
	"fmt"
)

type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return "Meow!"
}

func main() {
	dog := Dog{Name: "Buddy"}
	cat := Cat{Name: "Whiskers"}

	animals := []Speaker{dog, cat}

	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
