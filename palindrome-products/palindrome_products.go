package palindrome

import "fmt"

// NOTE: in this variant the benchmark increased speed by 15 times (20 vs 300)

// Product represets a palindrome.
type Product struct {
	Product        int
	Factorizations [][2]int
}

// Products finds the largest and smallest palindromes
// which are products of numbers within given range.
func Products(fmin, fmax int) (pmin, pmax Product, err error) {
	pmin.Product = fmin * fmax

	if fmax < fmin {
		err = fmt.Errorf("fmin > fmax")
		return
	}

	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			product := i * j

			// find reverse
			number, reverse := product, 0
			for number != 0 {
				reverse = reverse*10 + number%10
				number /= 10
			}

			if product == reverse && product >= pmax.Product {
				updateProduct(&pmax, product, i, j)
			}

			if product == reverse && product <= pmin.Product {
				updateProduct(&pmin, product, i, j)
			}
		}
	}

	if len(pmin.Factorizations) == 0 && len(pmax.Factorizations) == 0 {
		err = fmt.Errorf("no palindromes")
		return
	}

	return
}

// update Product struct
func updateProduct(p *Product, product, i, j int) {
	if product == p.Product {
		// only add new pair i and j
		p.Factorizations = append(p.Factorizations, [2]int{i, j})
	} else {
		p.Product = product
		p.Factorizations = [][2]int{{i, j}}
	}
}
