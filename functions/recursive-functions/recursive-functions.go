package main

import "fmt"

func factorial(n int) (int, error) {
	if n < 0 {
		return -1, fmt.Errorf("Número inválido: %d", n)
	}

	if n == 0 {
		return 1, nil
	}

	factorialBefore, _ := factorial(n - 1)
	return n * factorialBefore, nil
}

func factorialSimple(n uint) uint {
	if n == 0 {
		return 1
	}

	return n * factorialSimple(n-1)
}

func main() {
	result, _ := factorial(5)
	fmt.Println(result) // 120

	_, err := factorial(-4)

	if err != nil {
		fmt.Println(err) // Número inválido: -4
	}

	// result2 := factorialSimple(-4) --> error
	result2 := factorialSimple(4)
	fmt.Println(result2) // 24
}
