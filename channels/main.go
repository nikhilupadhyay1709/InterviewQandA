package main

import "fmt"

// Prints a hello message using values received in
// the channel
func hello(c chan string) {
	name := <-c
	fmt.Println("Hello", name)
}

func main() {

	m := make(chan string)

	go hello(m)
	m <- "World"
	close(m)

}