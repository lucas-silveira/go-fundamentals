package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println(reflect.TypeOf(10)) // int

	// números positivos uint8 uint16 uint32 uint64
	var b byte = 255               // byte é um alias para o tipo uint8
	fmt.Println(reflect.TypeOf(b)) // uint8

	// números negativos int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println(i1)                 // 9223372036854775807
	fmt.Println(reflect.TypeOf(i1)) // int

	var i2 rune = 'a'               // representa um mapeamento da tabela Unicode(int32)
	fmt.Println(reflect.TypeOf(i2)) // int32
	fmt.Println(i2)                 // 97

	// números reais (float32, float64)
	var x float32 = 49.99
	var x2 = float32(49.99)
	fmt.Println(reflect.TypeOf(x))     // float32
	fmt.Println(reflect.TypeOf(x2))    // float32
	fmt.Println(reflect.TypeOf(49.99)) // float64

	// boolean
	bo := true
	fmt.Println(reflect.TypeOf(bo)) // bool
	fmt.Println(!bo)                // false

	// string
	s1 := "Hello World"
	fmt.Println(reflect.TypeOf(s1)) // string
	fmt.Println(s1 + "!")           // Hello World!
	fmt.Println(len(s1))            // 11

	// string com múltiplas linhas
	s2 := `Hello
	World
	...
	`
	fmt.Println(len(s2)) // 19

	// char -> int32
	char := 'a'
	fmt.Println(reflect.TypeOf(char)) // int32
	fmt.Println(char)                 // 97
}
