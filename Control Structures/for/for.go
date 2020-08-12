package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1

	// em Go não existe o loop while
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d\n", i)
	}

	// loop infinito
	for {
		fmt.Println("Loop \"infinito\"...")
		time.Sleep(time.Second * 2) // A cada 2 segundos
	}
}
