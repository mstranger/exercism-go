package main

import "fmt"

type Product struct {
	Product        int
	Factorizations [][2]int
}

// Function to calculate largest
// palindrome which is product of
// two n-digits numbers
func palindrome(fmin, fmax int) (pmax, pmin Product) {

	upper_limit, lower_limit := fmax, fmin

	// initialize result
	// max_product := 0
	// for i := upper_limit; i >= lower_limit; i-- {
	// for j := i; j >= lower_limit; j-- {
	for i := lower_limit; i <= upper_limit; i++ {
		for j := i; j <= upper_limit; j++ {
			// calculating product of
			// two n-digit numbers
			product := i * j
			if product < pmax.Product {
				break
			}
			number := product
			reverse := 0

			// calculating reverse of
			// product to check whether
			// it is palindrome or not
			for number != 0 {
				reverse = reverse*10 + number%10
				number /= 10
			}

			// update new product if exist
			// and if greater than previous one
			if product == reverse && product >= pmax.Product {
				if product == pmax.Product {
					pmax.Factorizations = append(pmax.Factorizations, [2]int{i, j})
				} else {
					pmax.Product = product
					pmax.Factorizations = [][2]int{{i, j}}
				}

			}
		}
	}

	return
}

func main() {
	fmt.Println(palindrome(10, 99))
}
