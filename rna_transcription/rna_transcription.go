package rna_transcription

var dnaToRnaMap = map[string]string{
	"G": "C",
	"C": "G",
	"T": "A",
	"A": "U",
}

// ToRNA converts a DNA strand into an RNA strand
func ToRNA(dna string) string {
	result := ""
	for _, n := range dna {
		result += dnaToRnaMap[string(n)]
	}
	return result
}
