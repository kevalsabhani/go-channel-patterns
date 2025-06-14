package main

import "fmt"

func main() {

	ch := make(chan int)

	go func() {
		res := <-ch
		println("Received:", res)
	}()
	ch <- 42 // Send a value to the channel
	fmt.Println("Sent: 42")
}
