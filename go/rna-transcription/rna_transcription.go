package strand

import "strings"

var nucleotideComplement = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA converts a DNA strand to its RNA complement
func ToRNA(dna string) string {
	var rnaBuilder strings.Builder
	rnaBuilder.Grow(len(dna))

	for _, nucleotide := range dna {
		complement := nucleotideComplement[nucleotide]
		rnaBuilder.WriteRune(complement)
	}

	return rnaBuilder.String()
}
