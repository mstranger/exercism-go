package dna

import "fmt"

// Histogram type contains nucleotides and their count in DNA string.
type Histogram map[rune]int

// DNA string with nucleotides like "AAGGTC"
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
// Counts is a method on the DNA type.
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{
		'A': 0, 'C': 0, 'G': 0, 'T': 0,
	}
	for _, v := range d {
		if _, ok := h[v]; !ok {
			return nil, fmt.Errorf("error: invalid DNA")
		}

		h[v]++
	}
	return h, nil
}
