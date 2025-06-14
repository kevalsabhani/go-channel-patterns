package main

import (
	"fmt"
	"time"
)

func main() {
	const cap = 100
	ch := make(chan string, cap)
	go func() {
		for p := range ch {
			fmt.Println("child : recv'd signal :", p)
		}
	}()
	const work = 100
	for w := 0; w < work; w++ {
		select {
		case ch <- fmt.Sprintf("signal %d", w):
			fmt.Println("parent : sent signal :", w)
		default:
			fmt.Println("parent : dropped data :", w)
		}
	}
	close(ch)
	fmt.Println("parent : sent shutdown signal")
	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}
