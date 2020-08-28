package main

import (
	"fmt"
	"time"
)

func speak(person string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			c <- fmt.Sprintf("%s falando: %d", person, i)
		}
	}()

	return c
}

func join(channel1, channel2 <-chan string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			select {
			case result1 := <-channel1:
				channel <- result1
			case result2 := <-channel2:
				channel <- result2
			}
		}
	}()

	return channel
}

func main() {
	c := join(
		speak("John"),
		speak("Mary"),
	)

	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
}
