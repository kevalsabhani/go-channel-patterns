package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 100)
	sem := make(chan struct{}, 10) // Semaphore with a capacity of 10
	numGoroutines := 1000
	numMessages := 1000

	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			sem <- struct{}{}           // Acquire semaphore for the goroutine
			time.Sleep(1 * time.Second) // Simulate work
			fmt.Println("Goroutine", id, "is working")
			ch <- fmt.Sprintf("Goroutine %d", id) // Send message to channel
			<-sem
		}(i)
	}
	for i := 0; i < numMessages; i++ {
		fmt.Printf("Received message: %s\n", <-ch) // Receive messages from channel
	}
	close(ch)  // Close the channel after all messages are processed
	close(sem) // Close the semaphore channel (optional, for cleanup)
}
