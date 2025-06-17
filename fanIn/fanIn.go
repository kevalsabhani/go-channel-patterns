// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	chOne := channelGenerator("ChOne")
	chTwo := channelGenerator("ChTwo")

	ch := fanIn(chOne, chTwo)

	for i := 0; i < 20; i++ {
		fmt.Println(<-ch)
	}

}

func channelGenerator(channelName string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("Message %v from Channel %v", i, channelName)
		}
	}()
	return ch
}

func fanIn(childChOne, childChTwo <-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			ch <- <-childChOne
		}
	}()
	go func() {
		for {
			ch <- <-childChTwo
		}
	}()
	return ch
}
