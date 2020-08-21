package main

import "fmt"

func getApprovedValue(number int) int {
	defer fmt.Println("Fim") // adia a execução da função

	if number >= 5000 {
		fmt.Println("Valor alto...")
		return 5000
	}

	fmt.Println("Valor baixo...")
	return number
}

func main() {
	fmt.Println(getApprovedValue(6000))
	fmt.Println(getApprovedValue(4000))
}
