package main

import "fmt"

func main() {
	// var approveds map[int]string -> erro
	// maps devem ser inicializados

	approveds := make(map[int]string)

	approveds[123456789] = "Mark"
	approveds[987654321] = "John"
	approveds[321654987] = "Alice"

	fmt.Println(approveds) // map[123456789:Mark 321654987:Alice 987654321:John]

	for cpf, name := range approveds {
		fmt.Printf("%s (CPF: %d)\n", name, cpf)
	}

	fmt.Println(approveds[123456789]) // Mark

	delete(approveds, 123456789)

	fmt.Println(approveds) // map[321654987:Alice 987654321:John]
}
