package main

import (
	"runtime"
	"sync"
)

func main() {
	ch := make(chan int, 5)
	wg := sync.WaitGroup{}

	numGoroutines := runtime.GOMAXPROCS(0) // Number of goroutines to receive from the channel
	numMessages := 1000                    // Number of messages to send
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go func(id int, wg *sync.WaitGroup) {
			defer wg.Done() // Decrement the counter when the goroutine completes
			for num := range ch {
				println("Goroutine", id, "received:", num)
			}
		}(i, &wg)
	}

	for i := 0; i < numMessages; i++ {
		ch <- i
	}
	close(ch) // Close the channel when done
	wg.Wait() // Wait for all goroutines to finish
}
