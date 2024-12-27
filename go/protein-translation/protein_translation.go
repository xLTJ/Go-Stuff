package protein

import (
	"errors"
)

const codonSize = 3

var (
	// ErrStop is returned when a stop codon is encountered.
	ErrStop = errors.New("stop codon encountered")

	// ErrInvalidBase is returned when an invalid codon is encountered.
	ErrInvalidBase = errors.New("invalid or unknown codon encountered")
)

var codonToProteinMap = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

// FromRNA converts an RNA sequence into a list of proteins.
func FromRNA(rnaSequence string) ([]string, error) {
	var proteins []string
	for _, codon := range SplitIntoCodons(rnaSequence) {
		protein, err := FromCodon(codon)
		if err != nil {
			if errors.Is(err, ErrStop) {
				break
			}
			return proteins, ErrInvalidBase
		}
		proteins = append(proteins, protein)
	}
	return proteins, nil
}

// FromCodon translates a codon into the corresponding protein or returns an error.
func FromCodon(codon string) (string, error) {
	protein, exists := codonToProteinMap[codon]
	if !exists {
		return "", ErrInvalidBase
	}
	if protein == "STOP" {
		return "", ErrStop
	}
	return protein, nil
}

// SplitIntoCodons splits an RNA sequence into codons of fixed size.
func SplitIntoCodons(rnaSequence string) []string {
	codons := make([]string, 0, len(rnaSequence)/codonSize)
	for i := 0; i+codonSize <= len(rnaSequence); i += codonSize {
		codons = append(codons, rnaSequence[i:i+codonSize])
	}
	return codons
}
