// waitForResult.go
// waitForResult demonstrates how to wait for a result from a worker goroutine using channels.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Result represents the output of a work operation
type Result struct {
	Value string
	Error error
}

func main() {
	// Create buffered channel to prevent goroutine leak
	resCh := make(chan Result, 1)

	// Launch worker goroutine
	go doWork(resCh)

	// Wait for and handle result
	result := <-resCh
	if result.Error != nil {
		fmt.Printf("[Error] Work failed: %v\n", result.Error)
		return
	}
	fmt.Printf("[Info] Result received in main: %s\n", result.Value)
}

// doWork simulates a worker that performs some work and sends the result to the provided channel.
func doWork(resCh chan<- Result) {
	fmt.Println("[Info] Starting work...")

	// Simulate work with random duration
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

	// Send result through channel
	resCh <- Result{
		Value: "some result",
		Error: nil,
	}
	fmt.Println("[Info] Result sent from worker")
}
