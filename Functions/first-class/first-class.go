package main

import "fmt"

var sum = func(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(sum(2, 3)) // 5

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(3, 2)) // 1
}
