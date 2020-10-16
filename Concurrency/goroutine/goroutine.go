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
	// speak("Mary", "Pq vc não fala comigo?", 3)

	// o comando go cria uma nova corotina e executa a função de forma concorrente/paralela
	// go speak("Mary", "Ei...", 10)
	// go speak("John", "Opa...", 10)

	go speak("Mary", "Ei...", 10)
	speak("John", "Opa...", 5)
}
