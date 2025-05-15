package main

import (
	"fmt"
	"time"
)

// RateLimiter implements a simple token bucket algorithm.
type RateLimiter struct {
	tokens   chan struct{}
	interval time.Duration
}

// NewRateLimiter creates a new rate limiter that allows 'rps' requests per second.
func NewRateLimiter(rps int) *RateLimiter {
	rl := &RateLimiter{
		tokens:   make(chan struct{}, rps),
		interval: time.Second / time.Duration(rps),
	}

	// Continuously fill the token bucket.
	go func() {
		ticker := time.NewTicker(rl.interval)
		defer ticker.Stop()
		for range ticker.C {
			select {
			case rl.tokens <- struct{}{}:
			default: // Bucket is full, discard token
			}
		}
	}()

	return rl
}

// Allow returns true if a token is available (i.e., request is allowed), otherwise false.
func (r *RateLimiter) Allow() bool {
	select {
	case <-r.tokens:
		return true
	default:
		return false
	}
}

func main() {
	// Create a limiter that allows 2 requests per second
	limiter := NewRateLimiter(2)

	for i := 1; i <= 10; i++ {
		if limiter.Allow() {
			fmt.Printf("Request #%d allowed at %s\n", i, time.Now().Format("15:04:05.000"))
		} else {
			fmt.Printf("Request #%d denied at %s (Rate limit exceeded)\n", i, time.Now().Format("15:04:05.000"))
		}
		time.Sleep(300 * time.Millisecond)
	}
}
