package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	chone := channelCreator("chone")
	chtwo := channelCreator("chtwo")
	for i := 0; i < 5; i++ {
		fmt.Println(<-chone)
		fmt.Println(<-chtwo)
	}
}

func channelCreator(channelName string) chan string {
	ch := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Duration(rand.Intn(700)) * time.Millisecond)
			ch <- fmt.Sprintf("From %v goroutine: msg #%v", channelName, i)
		}
	}()
	return ch
}
