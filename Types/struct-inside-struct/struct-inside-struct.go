package main

import "fmt"

type item struct {
	productID int
	amount    int
	price     float64
}

type order struct {
	userID int
	itens  []item
}

func (o order) getTotal() float64 {
	total := 0.0

	for _, item := range o.itens {
		total += item.price * float64(item.amount)
	}

	return total
}

func main() {
	order := order{
		userID: 1,
		itens: []item{
			item{productID: 1, amount: 2, price: 12.10},
			item{productID: 2, amount: 1, price: 23.49},
			item{3, 10, 4.00}, // podemos omitir as chaves dos atributos
		},
	}

	fmt.Printf("O valor total do pedido é R$ %.2f", order.getTotal())
}
