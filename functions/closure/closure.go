package main

import "fmt"

func closure() func() {
	x := 1
	var function = func() {
		fmt.Println(x)
	}

	return function
}

func closure2(x int) func() {
	return func() {
		fmt.Println(x)
	}
}

func main() {
	x := 20
	fmt.Println(x) // 20
	closure()()    // 1

	printClosure := closure()
	printClosure() // 1

	printClosure2 := closure2(10)
	printClosure2() // 10

}
