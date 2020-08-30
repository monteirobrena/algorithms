/*
 * Selection Sort
 */

package main

import "fmt"

type Product struct {
	Name  string
	Price float32
}

func main() {
	products := []Product{
		{
			Name:  "Lamborghini",
			Price: 1000000,
		}, {
			Name:  "Jipe",
			Price: 46000,
		}, {
			Name:  "Brasília",
			Price: 16000,
		}, {
			Name:  "Smart",
			Price: 46000,
		}, {
			Name:  "Fusca",
			Price: 17000,
		},
	}

	sort(products)

	fmt.Println(products)
}

func sort(products []Product) {
	for current := 0; current < (len(products) - 1); current++ {
		lowestPrice := findLowestPrice(products, current, (len(products) - 1))

		productCurrent := products[current]
		productCheaper := products[lowestPrice]

		products[current] = productCheaper
		products[lowestPrice] = productCurrent
	}
}

func findLowestPrice(products []Product, first int, last int) int {
	cheaper := first

	for current := first; current <= last; current++ {
		if products[current].Price < products[cheaper].Price {
			cheaper = current
		}
	}

	return cheaper
}
