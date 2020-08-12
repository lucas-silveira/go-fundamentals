package main

import "fmt"

func rateToConcept(n float64) string {
	var rate = int(n)
	// em Go uma estrutura switch termina quando o primeiro case é satisfeito
	switch rate {
	case 10:
		fallthrough // indica para o switch avançar para o próximo case
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}
}

func main() {
	fmt.Println(rateToConcept(9.8))
	fmt.Println(rateToConcept(6.9))
	fmt.Println(rateToConcept(2.1))
	fmt.Println(rateToConcept(10.9))
	fmt.Println(rateToConcept(11))
}
