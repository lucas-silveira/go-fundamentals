package main

import "fmt"

type sporty interface {
	activeTurbo()
}

type ferrari struct {
	model        string
	turboOn      bool
	currentSpeed int
}

func (f *ferrari) activeTurbo() {
	f.turboOn = true
}

func main() {
	car1 := ferrari{"F40", false, 0}
	car1.activeTurbo()

	var car2 sporty = &ferrari{"F40", false, 0} // quando usamos métodos de interface que usam ponteiros
	// precisamos atribuir à variável o endereço do struct
	car2.activeTurbo()

	fmt.Println(car1, car2) // {F40 true 0} &{F40 true 0}
}
