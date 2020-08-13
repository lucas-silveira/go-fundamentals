package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Slice não é um array. Ele apenas reflete um pedaço de um array.
	// Em teoria, o slice possui somente um tamanho e um ponteiro para um array.
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  // slice

	fmt.Println(a1, s1)                                 // [1 2 3] [1 2 3]
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1)) // [3]int []int

	a2 := [5]int{1, 2, 3, 4, 5}
	s2 := a2[1:3]

	fmt.Println(a2, s2) // [1 2 3 4 5] [2 3]

	s3 := a2[:2]

	fmt.Println(a2, s3) // [1 2 3 4 5] [1 2]

	s4 := s2[:1] // slice de um slice

	fmt.Println(s2, s4) // [2 3] [2]
}
