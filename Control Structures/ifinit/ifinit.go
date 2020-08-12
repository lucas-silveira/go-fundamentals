package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomNumber() int {
	s := rand.NewSource(time.Now().UnixNano()) // gera um número aleatório a partir de um nano segundo
	r := rand.New(s)

	return r.Intn(10) // gera um número aleatório com valor máximo de 10
}

func main() {
	if i := generateRandomNumber(); i > 5 { // a variável i só estará disponível dentro do bloco if/else
		fmt.Println("Ganhou!")
	} else {
		fmt.Println("Perdeu!")
	}
}
