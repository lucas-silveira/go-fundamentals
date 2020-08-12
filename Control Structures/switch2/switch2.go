package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	// quando não adicionamos um parâmetro no switch, ele será finalizado
	// quando encontrar o primeiro case que retornar true (bool)
	fmt.Println(t)
	switch {
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Boa tarde!")
	default:
		fmt.Println("Boa noite!")
	}
}
