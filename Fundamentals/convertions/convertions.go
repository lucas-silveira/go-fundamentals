package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	// fmt.Println(x / y) -> error
	fmt.Println(x / float64(y)) // 1.2

	rate := 6.9
	finalRate := int(rate)
	fmt.Println(finalRate) // 6

	fmt.Println("Teste " + string(97)) // Teste a

	// int para string
	fmt.Println("Teste " + strconv.Itoa(97)) // Teste 97

	// string para int
	num, _ := strconv.Atoi("123")
	// Em Go as funções podem retornar dois valores. No caso acima, a função
	// retorna o valor que buscamos e um segundo valor que representa um erro
	// caso ele ocorra.
	fmt.Println(num)       // 123
	fmt.Println(num - 122) // 1

	b, _ := strconv.ParseBool("true")
	// a variável b só será verdadeiro para os valores "true" e "1"

	if b {
		fmt.Println(b) // true
	}
}
