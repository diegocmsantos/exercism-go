package protein

import (
	"errors"
	"strings"
)

var codonMap = map[string]string{
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

//ErrStop Errortype returned when stop
var ErrStop error

//ErrInvalidBase Errortype returned if incorrect Codon
var ErrInvalidBase error

// FromCodon Translates input
func FromCodon(input string) (string, error) {
	for _, letter := range input {
		if !strings.ContainsRune("AUGC", letter) {
			return "", ErrInvalidBase
		}
	}
	if codon, ok := codonMap[input]; ok {
		if codon != "STOP" {
			return codon, nil
		} else if codon == "STOP" {
			return "", ErrStop
		}
	}
	return "", errors.New("")
}

//FromRNA Translates RNA to an array of Codons
func FromRNA(rna string) ([]string, error) {
	var result []string
	for i := 0; i < len(rna); i = i + 3 {
		codon := rna[i : i+3]
		protein, ok := codonMap[codon]
		if !ok {
			return result, ErrInvalidBase
		}
		if protein == "STOP" {
			return result, nil
		}
		result = append(result, protein)
	}
	return result, ErrStop
}