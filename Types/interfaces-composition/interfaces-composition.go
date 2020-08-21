package main

import "fmt"

type sporty interface {
	activeTurbo()
}

type eletric interface {
	loadBattery()
}

type sportyEletric interface {
	sporty
	eletric
}

type bmw7 struct{}

func (b bmw7) activeTurbo() {
	fmt.Println("Turbo...")
}

func (b bmw7) loadBattery() {
	fmt.Println("Carregando bateria...")
}

func main() {
	var b sportyEletric = bmw7{}
	b.activeTurbo() // Turbo...
	b.loadBattery() // Carregando bateria...
}
