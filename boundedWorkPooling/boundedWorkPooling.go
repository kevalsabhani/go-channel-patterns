package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var numGoroutines = runtime.GOMAXPROCS(0)
	var numTasks = 3000
	var comm = make(chan string, 50)
	var wg = sync.WaitGroup{}

	wg.Add(numGoroutines)
	// Spinning up all workers
	for i := 1; i <= numGoroutines; i++ {
		go func(workerId int) {
			defer wg.Done()
			for task := range comm {
				fmt.Printf("Worker %v : Performing task - %v\n", workerId, task)
			}
		}(i)
	}

	// Adding tasks
	for i := 1; i <= numTasks; i++ {
		comm <- fmt.Sprintf("Task #%v", i)
	}

	// closing channel
	close(comm)

	// waiting for all tasks to be completed
	wg.Wait()

	fmt.Println("All tasks are completed")

}
