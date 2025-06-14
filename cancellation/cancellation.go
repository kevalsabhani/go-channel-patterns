package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan string)
	ctx, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
	defer cancel()

	go func() {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		fmt.Println("Child goroutine sending message")
		ch <- "Hello, World!"
	}()

	select {
	case msg := <-ch:
		fmt.Println("Received message:", msg)
	case <-ctx.Done():
		fmt.Println("Context cancelled:", ctx.Err())
	}

	fmt.Println("Main function completed")

}
