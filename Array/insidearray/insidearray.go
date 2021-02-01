package main

import "fmt"

func main() {
	// demonstrando que os slices apontam para o mesmo array
	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 3)

	fmt.Println(s1, s2) // [0 0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0 0 1 2 3]

	s1[0] = 7

	fmt.Println(s1, s2) // [7 0 0 0 0 0 0 0 0 0] [7 0 0 0 0 0 0 0 0 0 1 2 3]
}
