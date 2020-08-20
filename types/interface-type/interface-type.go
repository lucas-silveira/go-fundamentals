package main

import "fmt"

type course struct {
	name string
}

func main() {
	var something interface{}

	fmt.Println(something) // <nil>

	something = 3
	fmt.Println(something) // 3

	type dynamic interface{}

	var something2 dynamic = "Hello"
	fmt.Println(something2) // Hello

	something2 = true
	fmt.Println(something2) // true

	something2 = course{"Golang..."}
	fmt.Println(something2) // {Golang...}
}
