package main

import "fmt"

func printResult(rate float64) {
	if rate >= 7 {
		fmt.Println("Aprovado com nota", rate)
	} else {
		fmt.Println("Reprovado com nota", rate)
	}
}

func main() {
	printResult(7.3) // Aprovado com nota 7.3
	printResult(5.1) // Reprovado com nota 5.1
}
