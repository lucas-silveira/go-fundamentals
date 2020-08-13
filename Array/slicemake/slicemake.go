package main

import "fmt"

func main() {
	// Podemos criar um slice sem fazer referência a nenhum array. Nesse caso,
	// a função make cria internamente um array ao qual o slice fará referência.
	s := make([]int, 10) // slice com 10 posições e array interno com 10
	s[9] = 12

	fmt.Println(s) // [0 0 0 0 0 0 0 0 0 12]

	s2 := make([]int, 10, 20) // slice com 10 posições e array interno com 20

	fmt.Println(s2, len(s2), cap(s2)) // [0 0 0 0 0 0 0 0 0 0] 10 20
	// a função cap pega a capacidade máxima do slice, em outras palavras,
	// pega o tamanho do array interno

	s2 = append(s2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0) // adiciona novos elementos ao slice

	fmt.Println(s2, len(s2), cap(s2)) // [0 0 0 0 0 0 0 0 0 0 1 2 3 4 5 6 7 8 9 0] 20 20

	s2 = append(s2, 1)

	fmt.Println(s2, len(s2), cap(s2)) //[0 0 0 0 0 0 0 0 0 0 1 2 3 4 5 6 7 8 9 0 1] 21 40
	// quando o slice ultrapassa a sua capacidade máxima (tamanho do array interno),
	// essa capacidade automaticamente dobra de tamanho
}
