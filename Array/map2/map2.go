package main

import "fmt"

func main() {
	employees := map[string]float64{
		"John Stwart":  11325.45,
		"Alice Capler": 15456.78,
		"Peter Park":   12000.0,
	}

	employees["Jack Junior"] = 13500.0

	delete(employees, "Unknown") // deletar um elemento que não existe não gera erros

	for name, salary := range employees {
		fmt.Println(name, salary)
	}
}
