package main

import (
	"fmt"
	"time"
)

func speak(person, text string, length int) {
	for i := 0; i < length; i++ {
		time.Sleep(time.Second) // parar por um segundo
		fmt.Printf("%s: %s (iteração %d)\n", person, text, i+1)
	}
}

func main() {
	// Execução serial
	// speak("Maria", "Pq vc não fala comigo?", 3)

	// o comando go cria uma nova rotina e executa a função em uma "thread" separada
	// go speak("Maria", "Ei...", 10)
	// go speak("João", "Opa...", 10)

	go speak("Maria", "Ei...", 10)
	speak("João", "Opa...", 5)
}
