package main

import "fmt"

var integers [10]int

func init() {
	// Por padrão, a função init() é a primeira a ser executada quando o arquivo é lido
	fmt.Println("Inicializando...")
}

func init() {
	// a principal tarefa da initfunção é inicializar variáveis ​​globais
	// que não podem ser inicializadas no contexto global. Por exemplo,
	// inicialização de uma matriz.

	for i := 0; i < 10; i++ {
		integers[i] = i
	}
}

func main() {
	fmt.Println("Main...")
	fmt.Println(integers)
}
