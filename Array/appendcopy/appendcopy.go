package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	// array1 = append(array1, 4, 5, 6) -> error
	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1) // [1 2 3] [4 5 6]

	slice2 := make([]int, 2)

	copy(slice2, slice1) // copiamos do slice1 a quantidade de elementos que o slice2 suporta,
	// a contar do primeiro elemento do slice1

	fmt.Println(slice2) // [4 5]
}
