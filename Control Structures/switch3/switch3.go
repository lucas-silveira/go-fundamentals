package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	// o tipo interface{} nos permite trabalhar com tipos mais genéricos
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return "tipo desconhecido"
	}

}

func main() {
	fmt.Println(tipo(2.3))        // real
	fmt.Println(tipo(1))          // inteiro
	fmt.Println(tipo("text"))     // string
	fmt.Println(tipo(func() {}))  // função
	fmt.Println(tipo(time.Now())) // tipo desconhecido
}
