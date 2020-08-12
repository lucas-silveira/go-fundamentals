package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Strings:", "Banana" == "Banana") // true
	fmt.Println("!=", 3 != 2)                     // true
	fmt.Println("<", 3 < 2)                       // false
	fmt.Println(">", 3 > 2)                       // true
	fmt.Println("<=", 3 <= 2)                     // false
	fmt.Println(">=", 3 >= 2)                     // true

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println(d1, d2)                 // 1969-12-31 21:00:00 -0300 -03 | 1969-12-31 21:00:00 -0300 -03
	fmt.Println("Datas:", d1 == d2)     // true
	fmt.Println("Datas:", d1.Equal(d2)) // true

	type Person struct {
		name string
	}

	p1 := Person{"John"}
	p2 := Person{"John"}
	fmt.Println("Pessoas:", p1 == p2) // True
}
