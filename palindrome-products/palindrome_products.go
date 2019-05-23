package palindrome

import "fmt"

// NOTE: using goroutines has indcreased benchmark loop run from 300 to 10000

// Product represets a palindrome.
type Product struct {
	Product        int
	Factorizations [][2]int
}

// Products finds the largest and smallest palindromes
// which are products of numbers within given range.
func Products(fmin, fmax int) (pmin, pmax Product, err error) {
	// errof if the interval is incorrect
	if fmax < fmin {
		err = fmt.Errorf("fmin > fmax")
		return
	}

	c := make(chan Product)
	go palindromes(fmin, fmax, palindromeMin, c)
	go palindromes(fmin, fmax, palindromeMax, c)
	pmin, pmax = <-c, <-c
	// we don't know who will be first
	if pmax.Product < pmin.Product {
		pmin, pmax = pmax, pmin
	}

	// error if palindromes not found
	if len(pmin.Factorizations) == 0 && len(pmax.Factorizations) == 0 {
		err = fmt.Errorf("no palindromes")
		return
	}

	return
}

// find max palindrome
func palindromeMax(fmin, fmax int) (pmax Product) {
	for i := fmax; i >= fmin; i-- {
		for j := i; j >= fmin; j-- {
			product := i * j
			if product < pmax.Product {
				break
			}

			// {j, i} because for loop is descending, and we want {1, 9}
			if product == reverseInt(product) {
				updateProduct(&pmax, product, j, i)
			}
		}
	}

	return
}

// find min palindrome
func palindromeMin(fmin, fmax int) (pmin Product) {
	pmin.Product = fmin * fmax

	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			product := i * j
			if product > pmin.Product {
				break
			}

			if product == reverseInt(product) {
				updateProduct(&pmin, product, i, j)
			}
		}
	}

	return
}

// find palindrome using function f and send the result to the given channel
func palindromes(fmin, fmax int, f func(int, int) Product, c chan Product) {
	p := f(fmin, fmax)
	c <- p
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

// reverse int number
func reverseInt(n int) (r int) {
	for n != 0 {
		r = r*10 + n%10
		n /= 10
	}
	return
}
