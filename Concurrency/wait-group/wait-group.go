package main

import (
	"fmt"
	"sync"
)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		waitGroup.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		waitGroup.Done()
	}()

	waitGroup.Wait()
}
