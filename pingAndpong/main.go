package main

import (
	"fmt"
	"time"
)

// Constants for the duration of the game and channel buffer size
const gameDuration = 5 * time.Second

func main() {
	// Creating channels for communication between goroutines
	ping := make(chan string, 1)
	pong := make(chan string, 1)

	// Start the ping and pong goroutines using functions
	startPingGoroutine(ping)
	startPongGoroutine(pong)

	// Start the timer to signal the end of the game
	done := make(chan bool)
	go func() {
		time.Sleep(gameDuration)
		done <- true
	}()

	// Main loop to print "ping" and "pong"
	for {
		select {
		case msg := <-ping:
			fmt.Println(msg)
		case msg := <-pong:
			fmt.Println(msg)
		case <-done:
			fmt.Println("You are the winner!")
			return
		}
	}
}

// startPingGoroutine starts a goroutine that sends "ping" messages to the ping channel
func startPingGoroutine(ping chan<- string) {
	go func() {
		for {
			ping <- "ping"
			time.Sleep(500 * time.Millisecond)
		}
	}()
}

// startPongGoroutine starts a goroutine that sends "pong" messages to the pong channel
func startPongGoroutine(pong chan<- string) {
	go func() {
		for {
			pong <- "pong"
			time.Sleep(500 * time.Millisecond)
		}
	}()
}
