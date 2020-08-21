package main

import "fmt"

func average(numbers ...float64) float64 {
	total := 0.0

	for _, num := range numbers {
		total += num
	}

	return total / float64(len(numbers))
}

func printApproveds(approveds ...string) {
	fmt.Println("Lista de aprovados")

	for i, approved := range approveds {
		fmt.Printf("%d) %s\n", i+1, approved)
	}
}

func main() {
	fmt.Println(average(1, 2, 3, 4)) // 2.5
	fmt.Println(average(10, 20, 30)) // 20

	approveds := []string{"Mary", "Peter", "John", "Alice"}
	printApproveds(approveds...) // essa operação de spread só funciona com slices
}
