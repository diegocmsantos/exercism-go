package isogram

import "unicode"

// IsIsogram returns true if a string input has no duplicate letter, false otherwise.
func IsIsogram(input string) bool {
	letterMap := make(map[rune]bool, len(input))
	for _, letter := range input {
		if letter == '-' || letter == ' ' {
			continue
		}
		lowerLetter := unicode.ToLower(letter)
		if letterMap[lowerLetter] {
			return false
		}
		letterMap[lowerLetter] = true
	}
	return true
}
