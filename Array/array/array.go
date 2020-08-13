package main

import "fmt"

func main() {
	// Os arrays em Go são homogêneos (os seus valores devem ser do mesmo tipo)
	// e estáticos (o tamanho não varia)

	var rates [3]float64

	fmt.Println(rates) // [0 0 0]

	rates[0], rates[1], rates[2] = 7.8, 4.3, 9.1

	fmt.Println(rates)      // [7.8 4.3 9.1]
	fmt.Println(len(rates)) // 3

	total := 0.0

	for i := 0; i < len(rates); i++ {
		total += rates[i]
	}

	media := total / float64(len(rates))

	fmt.Printf("Média %.2f\n", media) // Média 7.07
}
