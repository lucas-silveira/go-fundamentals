package main

import "fmt"

func doPurchase(work1, work2 bool) (bool, bool, bool) {
	purchaseTv50 := work1 && work2
	purchaseTv32 := work1 != work2 // Ou exclusivo (or)
	purchaseIceCream := work1 || work2

	return purchaseTv50, purchaseTv32, purchaseIceCream
}

func main() {
	tv50, tv32, iceCream := doPurchase(true, true)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t\n", tv50, tv32, iceCream, !iceCream)

	// unary operators (apenas postfix)
	x := 1
	y := 2
	z := false

	x++ // x += 1, x = x + 1
	y-- // y -= 1, y = y - 1

	fmt.Println(x, y, !z) // 2 1 true

	// ternary operators does not exists in Go.
	// Alternatively, we can use an if structure
}
