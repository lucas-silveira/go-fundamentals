package main

import (
	"fmt"
	"time"
)

func twoThreeForTimes(base int, ch chan int) {
	time.Sleep(time.Second)
	ch <- 2 * base // enviando dados ao canal

	time.Sleep(time.Second)
	ch <- 3 * base

	time.Sleep(time.Second)
	ch <- 4 * base
}

func main() {
	ch := make(chan int)

	go twoThreeForTimes(2, ch)

	a, b := <-ch, <-ch // recebendo dados do canal
	fmt.Println(a, b)  // 4 6

	fmt.Println(<-ch) // 8
}
