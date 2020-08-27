package main

import (
	"fmt"
	"time"
)

func isPrimo(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func primos(n int, c chan int) {
	init := 2

	for i := 0; i < n; i++ {
		for primo := init; ; primo++ {
			if isPrimo(primo) {
				c <- primo
				init = primo + 1
				time.Sleep(time.Millisecond * 100)
				break // sai do laço
			}
		}
	}

	close(c) // indica que o channel não irá mais receber dados
	// essa função evita deadlocks no programa
}

func main() {
	c := make(chan int, 30)
	go primos(cap(c), c) // cap(c) retorna o tamanho do buffer do canal

	for primo := range c { // o range c é como um loop infinito que fica solicitando valores ao chanel c até que ele seja fechado
		fmt.Printf("%d ", primo)
	}

	fmt.Println("Fim!")
}
