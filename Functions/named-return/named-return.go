package main

import "fmt"

func revert(p1, p2 int) (second, first int) {
	first = p1
	second = p2

	return // retorno limpo: não há a necessidade de retornar as variáveis explicitamente
}

func main() {
	fmt.Println(revert(1, 2)) // 2 1

	x, y := revert(1, 2)

	fmt.Println(x, y) // 2 1
}
