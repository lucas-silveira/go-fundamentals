package main

// Channel é a forma de comunicação entre goroutines
// Channel é um tipo

import "fmt"

func main() {
	ch := make(chan int, 1) // criando um channel de inteiros com um buffer contendo 1 posição

	ch <- 1 // enviando para o canal o valor 1 (escrita)

	<-ch // recebendo o valor do canal (leitura) e descartando a mensagem

	ch <- 2
	fmt.Println(<-ch) // 2
}
