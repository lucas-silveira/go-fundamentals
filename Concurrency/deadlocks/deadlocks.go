package main

import (
	"fmt"
	"time"
)

func routine(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // operação bloqueante
	fmt.Println("O número 1 foi lido pelo receptor")
}

func main() {
	c := make(chan int) // canal sem buffer

	go routine(c)

	fmt.Println(<-c)
	fmt.Println("O dado foi impresso no console")
	fmt.Println(<-c) // deadlock -> Não há mais mensagens para serem lidas.
}
