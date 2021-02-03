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
	c := make(chan int) // canal sem buffer, isto significa que ao enviarmos
	// uma mensagem à esse canal, a operação ficará bloqueada na rotina client,
	// até que alguém leia a mensagem no canal em outra rotina.

	go routine(c)

	fmt.Println(<-c) // operação bloqueante
	fmt.Println("O dado foi impresso no console")
	fmt.Println(<-c) // deadlock -> Não há mais mensagens para serem lidas.
}
