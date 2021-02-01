package main

import "fmt"

func init() {
	// Podemos ter mais de uma função init por arquivo ou pacote.
	// Quando temos mais de uma função init em diferentes arquivos,
	// a ordem de chamada da função init ocorre por ordem alfabética
	// dos nomes dos arquivos.
	fmt.Println("Inicializando 2...")
}
