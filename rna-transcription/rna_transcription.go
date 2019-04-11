package strand

var nucleotides = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA converts DNA strand to its RNA complement
func ToRNA(dna string) string {
	result := make([]rune, 0)
	for _, v := range dna {
		result = append(result, nucleotides[v])
	}

	return string(result)
}
