package main

import "fmt"

func main() {
	p1 := Point{x: 2.0, y: 2.0}
	p2 := Point{2.0, 4.0}

	fmt.Println(catetos(p1, p2))  // 0 2
	fmt.Println(Distance(p1, p2)) // 2
}
