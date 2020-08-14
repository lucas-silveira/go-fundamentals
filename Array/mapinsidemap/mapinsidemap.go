package main

import "fmt"

func main() {
	employeesByLetter := map[string]map[string]float64{
		"A": {
			"Alice Capler": 15456.78,
		},
		"J": {
			"John Stwart":  11325.45,
			"Julian Clark": 10000.0,
		},
		"P": {
			"Peter Park": 12000.0,
		},
	}

	fmt.Println(employeesByLetter) // map[A:map[Alice Capler:15456.78] J:map[John Stwart:11325.45 Julian Clark:10000] P:map[Peter Park:12000]]

	delete(employeesByLetter, "P")

	for letter, employee := range employeesByLetter {
		fmt.Println(letter, employee)
		for name, salary := range employee {
			fmt.Println(name, salary)
		}
	}
}
