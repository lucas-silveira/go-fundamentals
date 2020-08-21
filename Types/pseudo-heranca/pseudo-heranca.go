package main

import "fmt"

type car struct {
	name         string
	currentSpeed int
}

type ferrari struct {
	car // campo anônimo.
	// aqui estamos usando o método de "composição" para trazer os atributos
	// de car para dentro do struct ferrari
	turboOn bool
}

func main() {
	f := ferrari{}
	f.name = "F40"
	f.currentSpeed = 0
	f.turboOn = true

	fmt.Printf("A ferrari %s está com turbo ligado? %v\n", f.name, f.turboOn)
	// A ferrari F40 está com turbo ligado? true

	fmt.Println(f) // {{F40 0} true}
}
