package palindrome

import (
	"fmt"
	"strconv"
)

// Product represets a palindrome.
type Product struct {
	Product        int
	Factorizations [][2]int
}

// Products finds the largest and smallest palindromes
// which are products of numbers within given range.
func Products(fmin, fmax int) (pmin, pmax Product, err error) {
	if fmax < fmin {
		err = fmt.Errorf("fmin > fmax")
		return
	}

	pmin = Product{fmax * fmax, [][2]int{}}
	pmax = Product{fmin * fmin, [][2]int{}}
	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			if p := i * j; isPalindrome(p) {
				if p <= pmin.Product {
					if p == pmin.Product {
						pmin.Factorizations = append(pmin.Factorizations, [2]int{i, j})
					} else {
						pmin.Product = p
						pmin.Factorizations = [][2]int{{i, j}}
					}
				}
				if p >= pmax.Product {
					if p == pmax.Product {
						pmax.Factorizations = append(pmax.Factorizations, [2]int{i, j})
					} else {
						pmax.Product = p
						pmax.Factorizations = [][2]int{{i, j}}
					}
				}
			}
		}
	}

	if len(pmin.Factorizations) == 0 && len(pmax.Factorizations) == 0 {
		err = fmt.Errorf("no palindromes")
		return
	}

	return
}

// check if given number is palindrome
func isPalindrome(n int) bool {
	s := strconv.Itoa(n)
	return s == reverse(s)
}

// reverse string
func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}
