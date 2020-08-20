package main

import (
	"encoding/json"
	"fmt"
)

// Em Go, os atributos com letras mínusculas são "privados" e atributos com
// a primeira letra maiúscula são "públicos"
type product struct {
	// mapeando a struct para json
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct to json
	p1 := product{1, "Notebook", 1899.90, []string{"Promoção", "Eletrônico"}}
	p1Json, _ := json.Marshal(p1)

	fmt.Println(p1Json)         // slice de bytes
	fmt.Println(string(p1Json)) // {"id":1,"name":"Notebook","price":1899.9,"tags":["Promoção","Eletrônico"]}

	// json to struct
	var p2 product
	jsonString := `{"id":2,"name":"Caneta","price":8.90,"tags":["Papelaria", "Importados"]}`
	json.Unmarshal([]byte(jsonString), &p2)

	fmt.Println(p2)         // {2 Caneta 8.9 [Papelaria Importados]}
	fmt.Println(p2.Tags[1]) // Importados
}
