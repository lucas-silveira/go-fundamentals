package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma =>", a+b)       // 5
	fmt.Println("Subtração =>", a-b)  // 1
	fmt.Println("Divisão =>", a/b)    // 1
	fmt.Println("Multiplicação", a*b) // 6
	fmt.Println("Módulo", a%b)        // 1

	// bitwise
	fmt.Println("E =>", a&b)   // 2
	fmt.Println("Ou =>", a|b)  // 3
	fmt.Println("Xor =>", a^b) // 1

	c := 3.0
	d := 2.0

	fmt.Println("Maior =>", math.Max(float64(a), float64(b))) // 3
	fmt.Println("Menor =>", math.Min(c, d))                   // 2
	fmt.Println("Exponenciação =>", math.Pow(c, d))           // 9
}
