package main

import (
	"fmt"
	"reflect"
	"strings"
)

type person struct {
	name     string
	lastname string
}

// por padrão, a passagem do receiver é por valor
func (p person) getFullname() string {
	return p.name + " " + p.lastname
}

// para conseguirmos alterar o struct dentro do método, precisamos
// passar o receiver por referência (endereço/ponteiro)
func (p *person) setFullname(fullname string) {
	parts := strings.Split(fullname, " ")
	fmt.Println(reflect.TypeOf(parts)) // []string (slice)
	p.name = parts[0]
	p.lastname = parts[1]
}

func main() {
	p1 := person{"Peter", "Park"}
	fmt.Println(p1.getFullname()) // Peter Park

	p1.setFullname("John Snow")
	fmt.Println(p1.getFullname()) // John Snow
}
