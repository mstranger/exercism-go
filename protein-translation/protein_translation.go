package protein

import "errors"

type Codon string
type Protein string

// ErrStop represents error type if STOP found.
var ErrStop = errors.New("stop found")

// ErrInvalidBase represents error type if invalid base or something else.
var ErrInvalidBase = errors.New("invalid base")

// Dict represents all possible translations.
var Dict = map[Codon]Protein{
	"AUG": "Methionine",
	"UUU": "Phenylalanine", "UUC": "Phenylalanine",
	"UUA": "Leucine", "UUG": "Leucine",
	"UCU": "Serine", "UCC": "Serine", "UCA": "Serine", "UCG": "Serine",
	"UAU": "Tyrosine", "UAC": "Tyrosine",
	"UGU": "Cysteine", "UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP", "UAG": "STOP", "UGA": "STOP",
}

// FromCodon translates the given codon into coresponding protein.
func FromCodon(c string) (string, error) {
	v, ok := Dict[Codon(c)]

	if !ok {
		return "", ErrInvalidBase
	}

	if v == "STOP" {
		return "", ErrStop
	}

	return string(v), nil
}

// FromRNA translates RNA string into proteins.
func FromRNA(s string) ([]string, error) {
	result := make([]string, 0)
	for i := 0; i <= len(s)-3; i += 3 {
		v, err := FromCodon(s[i : i+3])
		if err != nil {
			if err == ErrStop {
				return result, nil
			}
			return result, err
		}
		result = append(result, v)
	}

	return result, nil
}
