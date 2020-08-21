package main

// Iremos utilizar o pacote "area" criado dentro do
// nosso workspace GOPATH (~/go/src)

import (
	"fmt"

	"github.com/lucas-silveira/area"
)

func main() {
	fmt.Println(area.Circ(6.0))      // 113.0976
	fmt.Println(area.Rect(5.0, 2.0)) // 10
	// fmt.Println(area._EqTriangle(5.0, 2.0)) --> error
}
