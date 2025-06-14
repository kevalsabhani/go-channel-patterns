package main

import "fmt"

func main() {
	// Create a channel to send integers
	intChannel := make(chan int)

	numGoroutines := 2  // Number of goroutines to receive from the channel
	numMessages := 1000 // Number of messages to send

	// Start a goroutine to send integers to the channel
	go func() {
		for i := 0; i < numMessages; i++ {
			intChannel <- i
		}
		close(intChannel) // Close the channel when done
	}()

	// Start multiple goroutines to receive from the channel
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			for num := range intChannel {
				println("Goroutine", id, "received:", num)
			}
		}(i)
	}

	// Wait for user input to prevent the main function from exiting immediately
	var input string
	println("Press Enter to exit...")
	fmt.Scanln(&input)
}
