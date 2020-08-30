package selection

import (
	product "algorithms/product"
)

/*
 * Selection sort
 */

func Sort(products []product.Product) {
	for current := 0; current < (len(products) - 1); current++ {
		lowestPrice := findLowestPrice(products, current, (len(products) - 1))

		productCurrent := products[current]
		productCheaper := products[lowestPrice]

		products[current] = productCheaper
		products[lowestPrice] = productCurrent
	}
}

func findLowestPrice(products []product.Product, first int, last int) int {
	cheaper := first

	for current := first; current <= last; current++ {
		if products[current].Price < products[cheaper].Price {
			cheaper = current
		}
	}

	return cheaper
}
