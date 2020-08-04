package scrabble

import (
	"strings"
)

var pointMap = map[string]int{
	"AEIOULNRST": 1,
	"DG":         2,
	"BCMP":       3,
	"FHVWY":      4,
	"K":          5,
	"JX":         8,
	"QZ":         10,
}

// Score receives a word, find the points to each letter and return the sum of the points
func Score(word string) int {
	var total int
	for _, letter := range word {
		for k, v := range pointMap {
			total += strings.Count(k, strings.ToUpper(string(letter))) * v
		}
	}

	return total
}
