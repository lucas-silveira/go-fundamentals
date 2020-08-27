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
	// o buffer do canal é usado para definir o número de dados/mensagens que o canal pode suportar
	// canais sem buffer suportam 1 dado/mensagem por default
	c := make(chan int) // canal sem buffer

	go routine(c)

	fmt.Println(<-c)
	fmt.Println("O dado foi impresso no console")
	fmt.Println(<-c) // deadlock -> Não há mais mensagens para serem lidas.
}
