package main

import "fmt"

func main() {
	// Ponteiro é uma estrutura de dado que armazena um endereço de memória.

	i := 1 // variáveis do tipo int ocupam 8 bytes na memória, equivalente a 64 bits.

	// Go não possui aritmética de ponteiros, porém conseguimos compartilhar o ponteiro
	// para que possamos ter várias váriaveis com a mesma referência de memória.

	var p *int = nil // nil significa nulo
	p = &i           // atribuindo o endereço de uma variável a um ponteiro

	fmt.Println(&i, p) // 0xc00001a148 0xc00001a148
	fmt.Println(*p)    // 1

	*p++

	fmt.Println(i) // 2
}
