package main

import "fmt"

func increment1(n int) {
	n++
}

func increment2(n *int) {
	*n++ // aqui, o * é usado para desreferenciar o ponteiro, ou seja,
	// obter acesso ao valor para o qual o ponteiro aponta
}

func main() {
	n := 1

	increment1(n)  // passagem por valor
	fmt.Println(n) // 1

	// & é usado para obter o endereço da variável
	increment2(&n) // passagem por referência
	fmt.Println(n) // 2
}
