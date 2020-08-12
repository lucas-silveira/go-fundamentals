package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.141516
	xs := fmt.Sprint(x) // converte o valor para string

	fmt.Println("O valor de x é " + xs)  // O valor de x é 3.141516
	fmt.Println("O valor de x é", x)     // O valor de x é 3.141516
	fmt.Printf("O valor de x é %.2f", x) // O valor de x é 3.141516

	a := 1
	b := 1.8888
	c := false
	d := "hello"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d) // 1 1.888800 1.9 false hello
	fmt.Printf("\n%v %v %v %v", a, b, c, d)         // 1 1.888800 1.9 false hello
}
