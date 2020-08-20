package main

import "fmt"

type rate float64 // criando um tipo personalizado rate usando como base o tipo float64

func (r rate) between(start, end float64) bool {
	return float64(r) >= start && float64(r) <= end
}

func getConceptRate(r rate) string {
	switch {
	case r.between(9.0, 10.0):
		return "A"
	case r.between(7.0, 8.99):
		return "B"
	case r.between(5.0, 7.99):
		return "C"
	case r.between(3.0, 4.99):
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(getConceptRate(9.8)) // A
	fmt.Println(getConceptRate(6.9)) // C
	fmt.Println(getConceptRate(2.1)) // E
}
