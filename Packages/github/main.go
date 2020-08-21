package main

import (
	"fmt"

	"github.com/lucas-silveira/goarea"
)

// Iremos utilizar o pacote "goarea" baixado do repositório no github

func main() {
	fmt.Println(goarea.Circ(6.0))      // 113.0976
	fmt.Println(goarea.Rect(5.0, 2.0)) // 10
}
