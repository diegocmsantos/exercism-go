package rna_transcription

import "strings"

var newReplacer = strings.NewReplacer(
	"G", "C",
	"C", "G",
	"T", "A",
	"A", "U",
)

// ToRNA converts a DNA strand into an RNA strand
func ToRNA(dna string) string {
	return newReplacer.Replace(dna)
}
