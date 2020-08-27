package main

import (
	"fmt"
	"time"
)

func routine(ch chan int) {
	fmt.Println("Executou!")
	ch <- 1
	// fmt.Println("Executou!")
	ch <- 2
	// fmt.Println("Executou!")
	ch <- 3
	// fmt.Println("Executou!")
	ch <- 4 // o buffer do canal está lotado, pois o canal em questão possui um buffer de 3 posições
	// portanto a partir daqui a operação estará bloqueada até o canal liberar espaço
	// para um novo dado
	ch <- 5
	ch <- 6
}

func main() {
	// o buffer do canal é usado para definir o número de dados/mensagens
	// que o canal pode suportar simultaneamente
	// canais sem buffer suportam 1 dado/mensagem por default
	ch := make(chan int, 3) // o canal possui um buffer com 3 posições
	go routine(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)
}
