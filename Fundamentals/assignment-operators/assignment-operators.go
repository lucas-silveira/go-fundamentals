package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 3 // inferência de tipo
	i += 3 // 6
	i -= 3 // 3
	i /= 2 // 1
	i *= 2 // 2
	i %= 2 // 0

	x, y := 1, 2
	fmt.Println(x, y)

	x, y = y, x
	fmt.Println(x, y)
}
