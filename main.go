package main

import (
	product "algorithms/product"
	"algorithms/selection"
	"fmt"
)

/*
 * Selection sort
 */
func main() {
	products := []product.Product{
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

	selection.Sort(products)

	fmt.Println(products)
}
