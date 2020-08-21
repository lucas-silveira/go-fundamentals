package main

import "fmt"

type printable interface {
	toString() string
}

type person struct {
	name     string
	lastname string
}

type product struct {
	name  string
	price float64
}

// interfaces são implementadas implicitamente
func (p person) toString() string {
	return p.name + " " + p.lastname
}

func (p product) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.name, p.price)
}

func print(x printable) {
	fmt.Println(x.toString())
}

func main() {
	var person1 printable = person{"John", "Snow"}
	product1 := product{"Notebook", 1998.90}

	print(person1)  // John Snow
	print(product1) // Notebook - R$ 1998.90

	fmt.Println(person1.toString())  // John Snow
	fmt.Println(product1.toString()) // Notebook - R$ 1998.90
}
