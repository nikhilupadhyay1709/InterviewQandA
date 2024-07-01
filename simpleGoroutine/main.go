package main

import (
	"fmt"
	"time"
)

func main() {
	go hello()
	time.Sleep(1 * time.Second)
	fmt.Println("Hello from, 'Main' function")

}

func hello() {
	fmt.Println("Hello from, 'Hello' function")
}
